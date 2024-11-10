package otelx

import (
	"context"
	"fmt"
	"time"

	"github.com/blackhorseya/godine/app/infra/configx"
	"github.com/blackhorseya/godine/pkg/contextx"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/otlp/otlpmetric/otlpmetricgrpc"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc"
	"go.opentelemetry.io/otel/propagation"
	sdkmetric "go.opentelemetry.io/otel/sdk/metric"
	"go.opentelemetry.io/otel/sdk/resource"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.4.0"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	// Tracer is the global tracer.
	Tracer = otel.Tracer("otelx")

	// Meter is the global meter.
	Meter = otel.Meter("otelx")
)

// SDK is the OpenTelemetry SDK.
type SDK struct {
	target      string
	serviceName string
}

// NewSDK creates a new OpenTelemetry SDK.
func NewSDK(app *configx.Application) (*SDK, func(), error) {
	ctx := contextx.Background()

	instance := &SDK{
		target:      app.OTel.Target,
		serviceName: app.Name,
	}

	clean, err := instance.setupOTelSDK(ctx)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to setup OpenTelemetry SDK: %w", err)
	}

	// Set global tracer and meter.
	Tracer = otel.Tracer(app.Name)
	Meter = otel.Meter(app.Name)

	return instance, clean, nil
}

func (x *SDK) setupOTelSDK(ctx contextx.Contextx) (func(), error) {
	if x.target == "" {
		ctx.Warn("OpenTelemetry is disabled")
		return func() {
			ctx.Debug("noop")
		}, nil
	}

	ctx.Info(
		"setting up OpenTelemetry SDK",
		zap.String("service_name", x.serviceName),
		zap.String("otlp", x.target),
	)

	var shutdownFuncs []func(context.Context) error

	res, err := resource.New(ctx, resource.WithAttributes(semconv.ServiceNameKey.String(x.serviceName)))
	if err != nil {
		ctx.Error("failed to create resource", zap.Error(err))
		return nil, err
	}

	conn, err := initConn(x.target)
	if err != nil {
		ctx.Error("failed to create gRPC client", zap.Error(err))
		return nil, err
	}

	tracerProvider, err := newTracer(ctx, res, conn)
	if err != nil {
		ctx.Error("failed to create the Jaeger exporter", zap.Error(err))
		return nil, err
	}
	shutdownFuncs = append(shutdownFuncs, tracerProvider.Shutdown)

	meterProvider, err := newMeter(ctx, res, conn)
	if err != nil {
		ctx.Error("failed to create the OTLP exporter", zap.Error(err))
		return nil, err
	}
	shutdownFuncs = append(shutdownFuncs, meterProvider.Shutdown)

	return func() {
		ctx.Info("shutting down OpenTelemetry SDK")
		for _, fn := range shutdownFuncs {
			_ = fn(ctx)
		}
	}, nil
}

func initConn(target string) (*grpc.ClientConn, error) {
	conn, err := grpc.NewClient(target, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, fmt.Errorf("failed to create gRPC client: %w", err)
	}

	return conn, nil
}

func newTracer(
	ctx context.Context,
	res *resource.Resource,
	conn *grpc.ClientConn,
) (*sdktrace.TracerProvider, error) {
	exporter, err := otlptracegrpc.New(ctx, otlptracegrpc.WithGRPCConn(conn))
	if err != nil {
		return nil, fmt.Errorf("failed to create the Jaeger exporter: %w", err)
	}

	processor := sdktrace.NewBatchSpanProcessor(exporter)
	provider := sdktrace.NewTracerProvider(
		sdktrace.WithSampler(sdktrace.AlwaysSample()),
		sdktrace.WithResource(res),
		sdktrace.WithSpanProcessor(processor),
	)
	otel.SetTracerProvider(provider)

	otel.SetTextMapPropagator(propagation.NewCompositeTextMapPropagator(
		propagation.TraceContext{},
		propagation.Baggage{},
	))

	return provider, nil
}

func newMeter(
	ctx context.Context,
	res *resource.Resource,
	conn *grpc.ClientConn,
) (p *sdkmetric.MeterProvider, err error) {
	exporter, err := otlpmetricgrpc.New(ctx, otlpmetricgrpc.WithGRPCConn(conn))
	if err != nil {
		return nil, fmt.Errorf("failed to create the OTLP exporter: %w", err)
	}

	provider := sdkmetric.NewMeterProvider(
		sdkmetric.WithReader(sdkmetric.NewPeriodicReader(exporter, sdkmetric.WithInterval(3*time.Second))),
		sdkmetric.WithResource(res),
	)
	otel.SetMeterProvider(provider)

	return provider, nil
}
