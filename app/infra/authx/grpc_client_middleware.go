package authx

import (
	"context"
	"fmt"

	"github.com/blackhorseya/godine/app/infra/otelx"
	userM "github.com/blackhorseya/godine/entity/domain/user/model"
	"github.com/blackhorseya/godine/pkg/contextx"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

// UnaryClientInterceptor is used to create a new grpc unary client interceptor
func (x *Authx) UnaryClientInterceptor() grpc.UnaryClientInterceptor {
	return func(
		c context.Context,
		method string,
		req, reply any,
		cc *grpc.ClientConn,
		invoker grpc.UnaryInvoker,
		opts ...grpc.CallOption,
	) error {
		next, span := otelx.Tracer.Start(c, "authx.grpc.UnaryClientInterceptor")
		defer span.End()

		if x.SkipPath(method) {
			return invoker(next, method, req, reply, cc, opts...)
		}

		ctx, err := contextx.FromContext(c)
		if err != nil {
			return fmt.Errorf("get context error: %w", err)
		}

		handler, err := userM.FromContext(c)
		if err != nil {
			ctx.Error("get user model from context error", zap.Error(err))
			return err
		}
		ctx.Debug("unary client interceptor", zap.Any("handler", &handler))

		c = metadata.NewOutgoingContext(next, metadata.New(map[string]string{
			"access_token": handler.AccessToken,
		}))

		return invoker(c, method, req, reply, cc, opts...)
	}
}

// StreamClientInterceptor is used to create a new grpc stream client interceptor
func (x *Authx) StreamClientInterceptor() grpc.StreamClientInterceptor {
	return func(
		c context.Context,
		desc *grpc.StreamDesc,
		cc *grpc.ClientConn,
		method string,
		streamer grpc.Streamer,
		opts ...grpc.CallOption,
	) (grpc.ClientStream, error) {
		ctx, err := contextx.FromContextLegacy(c)
		if err != nil {
			return nil, fmt.Errorf("get context error: %w", err)
		}

		ctx, span := otelx.Span(ctx, "authx.grpc.UnaryClientInterceptor")
		defer span.End()

		if x.SkipPath(method) {
			ctx.Debug("unary client interceptor", zap.String("method", method))
			return streamer(ctx, desc, cc, method, opts...)
		}

		handler, err := userM.FromContext(c)
		if err != nil {
			ctx.Error("get user model from context error", zap.Error(err))
			return nil, err
		}
		ctx.Debug("unary client interceptor", zap.Any("handler", &handler))

		c = metadata.NewOutgoingContext(c, metadata.New(map[string]string{
			"access_token": handler.AccessToken,
		}))

		return streamer(c, desc, cc, method, opts...)
	}
}
