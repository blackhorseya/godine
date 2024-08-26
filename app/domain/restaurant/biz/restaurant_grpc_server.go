package biz

import (
	"context"
	"fmt"
	"strconv"

	"github.com/blackhorseya/godine/app/infra/otelx"
	"github.com/blackhorseya/godine/entity/domain/restaurant/biz"
	"github.com/blackhorseya/godine/entity/domain/restaurant/model"
	"github.com/blackhorseya/godine/entity/domain/restaurant/repo"
	userM "github.com/blackhorseya/godine/entity/domain/user/model"
	"github.com/blackhorseya/godine/pkg/contextx"
	"go.uber.org/zap"
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

func (i *restaurantService) ListRestaurants(
	req *biz.ListRestaurantsRequest,
	stream biz.RestaurantService_ListRestaurantsServer,
) error {
	next, span := otelx.Tracer.Start(stream.Context(), "restaurant.biz.ListRestaurants")
	defer span.End()

	ctx, err := contextx.FromContext(stream.Context())
	if err != nil {
		return status.Errorf(codes.Internal, "failed to get contextx: %v", err)
	}

	items, total, err := i.restaurants.List(next, repo.ListCondition{
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
	ctx, err := contextx.FromContext(c)
	if err != nil {
		return nil, fmt.Errorf("failed to get contextx: %w", err)
	}

	ctx, span := otelx.Span(ctx, "restaurant.biz.GetRestaurant")
	defer span.End()

	return i.restaurants.GetByID(ctx, req.RestaurantId)
}
