package otelx

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/blackhorseya/godine/app/infra/configx"
	"github.com/blackhorseya/godine/pkg/contextx"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/otlp/otlpmetric/otlpmetricgrpc"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc"
	"go.opentelemetry.io/otel/metric"
	"go.opentelemetry.io/otel/propagation"
	sdkmetric "go.opentelemetry.io/otel/sdk/metric"
	"go.opentelemetry.io/otel/sdk/resource"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.4.0"
	"go.opentelemetry.io/otel/trace"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	// Tracer is the global tracer.
	Tracer = otel.Tracer("")

	// Meter is the global meter.
	Meter = otel.Meter("")
)

// OTelx is the OpenTelemetry SDK.
type OTelx struct {
	Tracer trace.Tracer
	Meter  metric.Meter

	target      string
	serviceName string
}

// New creates a new OpenTelemetry SDK.
func New(app *configx.Application) (*OTelx, func(), error) {
	ctx := contextx.Background()
	if app.OTel.Target == "" {
		ctx.Warn("OpenTelemetry is disabled")
		return &OTelx{
			Tracer:      otel.Tracer(app.Name),
			Meter:       otel.Meter(app.Name),
			target:      "",
			serviceName: app.Name,
		}, nil, nil
	}

	return &OTelx{
		target:      app.OTel.Target,
		serviceName: app.Name,
	}, func() {}, nil
}

// Shutdown is the function to shutdown the OpenTelemetry SDK.
var Shutdown = func(context.Context) error {
	return nil
}

// SetupOTelSDK sets up the OpenTelemetry SDK with the Jaeger exporter.
func SetupOTelSDK(ctx contextx.Contextx, app *configx.Application) (err error) {
	if app.OTel.Target == "" {
		ctx.Warn("OpenTelemetry is disabled")
		return nil
	}

	ctx.Info(
		"setting up OpenTelemetry SDK",
		zap.String("service_name", app.Name),
		zap.String("otlp", app.OTel.Target),
	)

	var shutdownFuncs []func(context.Context) error

	Shutdown = func(ctx context.Context) error {
		for _, fn := range shutdownFuncs {
			err = errors.Join(err, fn(ctx))
		}
		shutdownFuncs = nil
		return err
	}

	res, err := resource.New(ctx, resource.WithAttributes(semconv.ServiceNameKey.String(app.Name)))
	if err != nil {
		return fmt.Errorf("failed to create resource: %w", err)
	}

	conn, err := initConn(app)
	if err != nil {
		return err
	}

	tracerProvider, err := newTracer(ctx, res, conn, app)
	if err != nil {
		return err
	}
	shutdownFuncs = append(shutdownFuncs, tracerProvider.Shutdown)

	meterProvider, err := newMeter(ctx, res, conn, app)
	if err != nil {
		return err
	}
	shutdownFuncs = append(shutdownFuncs, meterProvider.Shutdown)

	return nil
}

func initConn(app *configx.Application) (*grpc.ClientConn, error) {
	conn, err := grpc.NewClient(app.OTel.Target, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, fmt.Errorf("failed to create gRPC client: %w", err)
	}

	return conn, nil
}

func newTracer(
	ctx context.Context,
	res *resource.Resource,
	conn *grpc.ClientConn,
	app *configx.Application,
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

	Tracer = provider.Tracer(app.Name)

	return provider, nil
}

func newMeter(
	ctx context.Context,
	res *resource.Resource,
	conn *grpc.ClientConn,
	app *configx.Application,
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

	Meter = provider.Meter(app.Name)

	return provider, nil
}
