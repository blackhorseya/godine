//go:build integration

package restaurant

import (
	"context"
	"net"
	"testing"

	"github.com/blackhorseya/godine/app/infra/storage/mongodbx"
	"github.com/blackhorseya/godine/app/infra/storage/redix"
	"github.com/blackhorseya/godine/entity/domain/restaurant/biz"
	"github.com/blackhorseya/godine/pkg/contextx"
	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_recovery "github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	grpc_ctxtags "github.com/grpc-ecosystem/go-grpc-middleware/tags"
	"github.com/redis/go-redis/v9"
	"github.com/stretchr/testify/suite"
	"go.mongodb.org/mongo-driver/mongo"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/resolver"
	"google.golang.org/grpc/test/bufconn"
)

type suiteRestaurantServiceIntegration struct {
	suite.Suite

	mongodbContainer *mongodbx.Container
	rw               *mongo.Client

	redisContainer *redix.Container
	rdb            *redis.Client

	baseServer *grpc.Server
	server     biz.RestaurantServiceServer
	client     biz.RestaurantServiceClient
}

func (s *suiteRestaurantServiceIntegration) SetupTest() {
	mongodbContainer, err := mongodbx.NewContainer(contextx.Background())
	s.Require().NoError(err)
	s.mongodbContainer = mongodbContainer

	rw, err := mongodbContainer.RW(contextx.Background())
	s.Require().NoError(err)
	s.rw = rw

	redisContainer, err := redix.NewContainer(contextx.Background())
	s.Require().NoError(err)
	s.redisContainer = redisContainer

	rdb, err := redisContainer.RW(contextx.Background())
	s.Require().NoError(err)
	s.rdb = rdb

	server, err := NewIntegration(s.rw, s.rdb)
	s.Require().NoError(err)
	s.server = server

	buffer := 10 * 1024 * 1024
	listen := bufconn.Listen(buffer)
	s.baseServer = grpc.NewServer(
		grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(
			grpc_ctxtags.UnaryServerInterceptor(),
			grpc_recovery.UnaryServerInterceptor(),
		)),
		grpc.StreamInterceptor(grpc_middleware.ChainStreamServer(
			grpc_ctxtags.StreamServerInterceptor(),
			grpc_recovery.StreamServerInterceptor(),
		)),
	)
	biz.RegisterRestaurantServiceServer(s.baseServer, s.server)
	go func() {
		if err = s.baseServer.Serve(listen); err != nil {
			s.T().Fatalf("failed to serve: %v", err)
		}
	}()

	resolver.SetDefaultScheme("passthrough")
	conn, err := grpc.NewClient(
		"bufnet",
		grpc.WithContextDialer(func(c context.Context, s string) (net.Conn, error) {
			return listen.Dial()
		}),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	s.Require().NoError(err)

	s.client = biz.NewRestaurantServiceClient(conn)
}

func (s *suiteRestaurantServiceIntegration) TearDownTest() {
	if s.baseServer != nil {
		s.baseServer.Stop()
	}

	if s.rw != nil {
		_ = s.rw.Disconnect(contextx.Background())
	}

	if s.rdb != nil {
		_ = s.rdb.Close()
	}

	if s.mongodbContainer != nil {
		_ = s.mongodbContainer.Terminate(contextx.Background())
	}

	if s.redisContainer != nil {
		_ = s.redisContainer.Terminate(contextx.Background())
	}
}

func TestIntegrationAll(t *testing.T) {
	suite.Run(t, new(suiteRestaurantServiceIntegration))
}

func (s *suiteRestaurantServiceIntegration) Test_restaurantService_ListRestaurants() {
	type args struct {
		c    context.Context
		req  *biz.ListRestaurantsRequest
		mock func()
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "list restaurants",
			args: args{
				c: context.TODO(),
				req: &biz.ListRestaurantsRequest{
					Page:     0,
					PageSize: 0,
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		s.T().Run(tt.name, func(t *testing.T) {
			if tt.args.mock != nil {
				tt.args.mock()
			}

			_, err := s.client.ListRestaurants(tt.args.c, tt.args.req)
			if tt.wantErr {
				s.Require().Error(err)
				return
			}

			s.Require().NoError(err)
		})
	}
}
