package platform

import (
	"github.com/blackhorseya/godine/app/infra/transports/grpcx"
	opsB "github.com/blackhorseya/godine/entity/domain/logistics/biz"
	notifyB "github.com/blackhorseya/godine/entity/domain/notification/biz"
	orderB "github.com/blackhorseya/godine/entity/domain/order/biz"
	payB "github.com/blackhorseya/godine/entity/domain/payment/biz"
	restB "github.com/blackhorseya/godine/entity/domain/restaurant/biz"
	userB "github.com/blackhorseya/godine/entity/domain/user/biz"
	"google.golang.org/grpc"
	"google.golang.org/grpc/health"
	"google.golang.org/grpc/health/grpc_health_v1"
	"google.golang.org/grpc/reflection"
)

// NewInitServersFn creates and returns a new InitServers function.
func NewInitServersFn(
	accountServer userB.AccountServiceServer,
	restaurantServer restB.RestaurantServiceServer,
	menuServer restB.MenuServiceServer,
	paymentServer payB.PaymentServiceServer,
	notifyServer notifyB.NotificationServiceServer,
	orderServer orderB.OrderServiceServer,
	logisticsServer opsB.LogisticsServiceServer,
) grpcx.InitServers {
	return func(s *grpc.Server) {
		// register health server
		healthServer := health.NewServer()
		grpc_health_v1.RegisterHealthServer(s, healthServer)
		healthServer.SetServingStatus(serviceName, grpc_health_v1.HealthCheckResponse_SERVING)

		// register reflection service on gRPC server.
		reflection.Register(s)

		// register services
		userB.RegisterAccountServiceServer(s, accountServer)
		restB.RegisterRestaurantServiceServer(s, restaurantServer)
		restB.RegisterMenuServiceServer(s, menuServer)
		payB.RegisterPaymentServiceServer(s, paymentServer)
		notifyB.RegisterNotificationServiceServer(s, notifyServer)
		orderB.RegisterOrderServiceServer(s, orderServer)
		opsB.RegisterLogisticsServiceServer(s, logisticsServer)
	}
}
