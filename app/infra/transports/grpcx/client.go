package grpcx

import (
	"fmt"

	"github.com/blackhorseya/godine/app/infra/configx"
	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	"go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// Client is the grpc client
type Client struct {
	services map[string]*configx.Application
	options  []grpc.DialOption
}

// NewClient is used to create a new grpc client
func NewClient(config *configx.Configuration) (*Client, error) {
	options := []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithStatsHandler(otelgrpc.NewClientHandler()),
		grpc.WithUnaryInterceptor(grpc_middleware.ChainUnaryClient()),
		grpc.WithStreamInterceptor(grpc_middleware.ChainStreamClient()),
	}

	return &Client{
		services: config.Services,
		options:  options,
	}, nil
}

// Dial is used to dial the grpc service
func (c *Client) Dial(service string) (*grpc.ClientConn, error) {
	app, ok := c.services[service]
	if !ok {
		return nil, fmt.Errorf("service: [%s] not found", service)
	}

	target := fmt.Sprintf("%s:%d", app.GRPC.URL, app.GRPC.Port)

	return grpc.NewClient(target, c.options...)
}
