package authx

import (
	"context"

	"github.com/blackhorseya/godine/app/infra/otelx"
	"github.com/blackhorseya/godine/entity/domain/user/model"
	"github.com/blackhorseya/godine/pkg/contextx"
	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

const keyAccessToken = "access_token"

// UnaryServerInterceptor is used to create a new unary interceptor
func (x *Authx) UnaryServerInterceptor() grpc.UnaryServerInterceptor {
	return func(c context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp any, err error) {
		next, span := otelx.Tracer.Start(c, "authx.grpc.UnaryServerInterceptor")
		defer span.End()

		ctx, err := contextx.FromContext(c)
		if err != nil {
			return nil, status.Newf(codes.Internal, "failed to get contextx: %v", err).Err()
		}

		if x.SkipPath(info.FullMethod) {
			ctx.Debug(
				"skip authx middleware",
				zap.Strings("skip_paths", x.SkipPaths),
				zap.String("full_method", info.FullMethod),
			)
			return handler(c, req)
		}

		account, err := extractAccount(c, x)
		if err != nil {
			ctx.Error("failed to extract account", zap.Error(err))
			return nil, err
		}
		next = context.WithValue(next, contextx.KeyHandler, account)

		return handler(next, req)
	}
}

// StreamServerInterceptor is used to create a new stream interceptor
func (x *Authx) StreamServerInterceptor() grpc.StreamServerInterceptor {
	return func(srv any, stream grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
		next, span := otelx.Tracer.Start(stream.Context(), "authx.grpc.StreamServerInterceptor")
		defer span.End()

		ctx, err := contextx.FromContext(stream.Context())
		if err != nil {
			return status.Errorf(codes.Internal, "failed to get contextx: %v", err)
		}

		if x.SkipPath(info.FullMethod) {
			ctx.Debug(
				"skip authx middleware",
				zap.Strings("skip_paths", x.SkipPaths),
				zap.String("full_method", info.FullMethod),
			)
			return handler(srv, stream)
		}

		account, err := extractAccount(stream.Context(), x)
		if err != nil {
			ctx.Error("failed to extract account", zap.Error(err))
			return err
		}
		next = context.WithValue(next, contextx.KeyHandler, account)

		wrappedStream := grpc_middleware.WrapServerStream(stream)
		wrappedStream.WrappedContext = next

		return handler(srv, stream)
	}
}

func extractAccount(c context.Context, authx *Authx) (*model.Account, error) {
	headers, ok := metadata.FromIncomingContext(c)
	if !ok {
		return nil, status.Errorf(codes.Unauthenticated, "metadata not found")
	}

	tokens := headers.Get(keyAccessToken)
	if len(tokens) < 1 {
		return nil, status.Errorf(codes.Unauthenticated, "access token not found")
	}
	accessToken := tokens[0]

	account, err := authx.ExtractAccountFromToken(accessToken)
	if err != nil {
		return nil, status.Errorf(codes.Unauthenticated, "failed to extract account: %v", err)
	}

	return account, nil
}
