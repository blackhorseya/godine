package domain

import (
	"context"
	"strconv"

	"github.com/blackhorseya/godine/app/infra/otelx"
	"github.com/blackhorseya/godine/entity/domain/restaurant/biz"
	"github.com/blackhorseya/godine/entity/domain/restaurant/model"
	"github.com/blackhorseya/godine/entity/domain/restaurant/repo"
	userM "github.com/blackhorseya/godine/entity/domain/user/model"
	"github.com/blackhorseya/godine/pkg/contextx"
	"github.com/blackhorseya/godine/pkg/persistence"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

type restaurantService struct {
	restaurants repo.IRestaurantRepo
}

// NewRestaurantService will create a new restaurant service.
func NewRestaurantService(restaurants repo.IRestaurantRepo) biz.RestaurantServiceServer {
	return &restaurantService{
		restaurants: restaurants,
	}
}

func (i *restaurantService) CreateRestaurant(
	c context.Context,
	req *biz.CreateRestaurantRequest,
) (*model.Restaurant, error) {
	next, span := otelx.Tracer.Start(c, "restaurant.biz.CreateRestaurant")
	defer span.End()

	ctx := contextx.WithContextx(c)

	handler, err := userM.FromContext(c)
	if err != nil {
		ctx.Error("failed to get user from context", zap.Error(err))
		return nil, status.Error(codes.Unauthenticated, err.Error())
	}

	restaurant := model.NewRestaurant(req.Name, req.Address)
	restaurant.CreatedBy = handler.Id

	err = i.restaurants.Create(next, restaurant)
	if err != nil {
		ctx.Error("failed to create restaurant", zap.Error(err))
		return nil, status.Error(codes.Internal, err.Error())
	}

	return restaurant, nil
}

func (i *restaurantService) ListRestaurants(
	req *biz.ListRestaurantsRequest,
	stream biz.RestaurantService_ListRestaurantsServer,
) error {
	c := stream.Context()
	next, span := otelx.Tracer.Start(c, "restaurant.biz.ListRestaurants")
	defer span.End()

	ctx := contextx.WithContextx(c)

	items, total, err := i.restaurants.List(next, persistence.Pagination{
		Limit:  req.PageSize,
		Offset: (req.Page - 1) * req.PageSize,
	})
	if err != nil {
		ctx.Error("failed to list restaurants", zap.Error(err))
		return status.Error(codes.Internal, err.Error())
	}

	err = stream.SetHeader(metadata.New(map[string]string{"total": strconv.Itoa(total)}))
	if err != nil {
		ctx.Error("failed to set header", zap.Error(err))
		return status.Error(codes.Internal, err.Error())
	}

	for _, item := range items {
		err = stream.Send(item)
		if err != nil {
			ctx.Error("failed to send restaurant", zap.Error(err))
			return status.Error(codes.Internal, err.Error())
		}
	}

	return nil
}

func (i *restaurantService) GetRestaurant(
	c context.Context,
	req *biz.GetRestaurantRequest,
) (*model.Restaurant, error) {
	next, span := otelx.Tracer.Start(c, "restaurant.biz.GetRestaurant")
	defer span.End()

	return i.restaurants.GetByID(next, req.RestaurantId)
}

func (i *restaurantService) ListRestaurantsNonStream(
	c context.Context,
	req *biz.ListRestaurantsRequest,
) (*biz.ListRestaurantsResponse, error) {
	next, span := otelx.Tracer.Start(c, "restaurant.biz.ListRestaurantsNonStream")
	defer span.End()

	ctx := contextx.WithContextx(c)

	items, total, err := i.restaurants.List(next, persistence.Pagination{
		Limit:  req.PageSize,
		Offset: (req.Page - 1) * req.PageSize,
	})
	if err != nil {
		ctx.Error("failed to list restaurants", zap.Error(err))
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &biz.ListRestaurantsResponse{
		Restaurants: items,
		Total:       int64(total),
	}, nil
}

func (i *restaurantService) PlaceOrder(c context.Context, req *biz.PlaceOrderRequest) (*biz.PlaceOrderResponse, error) {
	// TODO: 2024/10/2|sean|implement me
	panic("implement me")
}

func (i *restaurantService) ListOrders(
	req *biz.ListOrdersRequest,
	stream grpc.ServerStreamingServer[model.Order],
) error {
	// TODO: 2024/10/2|sean|implement me
	panic("implement me")
}
