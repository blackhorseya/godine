package contextx

import (
	"context"

	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	"google.golang.org/grpc"
)

// UnaryServerInterceptor is used to create a new unary interceptor
func UnaryServerInterceptor() grpc.UnaryServerInterceptor {
	return func(
		c context.Context,
		req any,
		info *grpc.UnaryServerInfo,
		handler grpc.UnaryHandler,
	) (resp any, err error) {
		ctx := WithContext(c)

		return handler(ctx, req)
	}
}

// StreamServerInterceptor is used to create a new stream interceptor
func StreamServerInterceptor() grpc.StreamServerInterceptor {
	return func(srv any, ss grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
		ctx := WithContext(ss.Context())
		wrappedStream := grpc_middleware.WrapServerStream(ss)
		wrappedStream.WrappedContext = ctx

		return handler(srv, ss)
	}
}
