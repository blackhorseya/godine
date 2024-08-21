package biz

import (
	"context"
	"fmt"

	"github.com/blackhorseya/godine/app/infra/otelx"
	"github.com/blackhorseya/godine/entity/domain/restaurant/biz"
	"github.com/blackhorseya/godine/entity/domain/restaurant/model"
	"github.com/blackhorseya/godine/entity/domain/restaurant/repo"
	userM "github.com/blackhorseya/godine/entity/domain/user/model"
	"github.com/blackhorseya/godine/pkg/contextx"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
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
	ctx, err := contextx.FromContext(c)
	if err != nil {
		return nil, fmt.Errorf("failed to get contextx: %w", err)
	}

	ctx, span := otelx.Span(ctx, "restaurant.biz.CreateRestaurant")
	defer span.End()

	handler, err := userM.FromContext(ctx)
	if err != nil {
		ctx.Error("failed to get user from context", zap.Error(err))
		return nil, status.Error(codes.Unauthenticated, err.Error())
	}

	restaurant := model.NewRestaurant(req.Name, req.Address)
	restaurant.CreatedBy = handler.Id

	err = i.restaurants.Create(ctx, restaurant)
	if err != nil {
		ctx.Error("failed to create restaurant", zap.Error(err))
		return nil, status.Error(codes.Internal, err.Error())
	}

	return restaurant, nil
}

func (i *restaurantService) GetRestaurant(
	ctx context.Context,
	request *biz.GetRestaurantRequest,
) (*model.Restaurant, error) {
	// TODO: 2024/8/21|sean|implement me
	panic("implement me")
}

func (i *restaurantService) ListRestaurants(
	request *biz.ListRestaurantsRequest,
	stream biz.RestaurantService_ListRestaurantsServer,
) error {
	// TODO: 2024/8/21|sean|implement me
	panic("implement me")
}
