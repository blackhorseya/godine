package handlers

import (
	"context"

	"connectrpc.com/connect"
	"github.com/blackhorseya/godine/entity/domain/restaurant/biz"
	"github.com/blackhorseya/godine/entity/domain/restaurant/biz/bizconnect"
	"github.com/blackhorseya/godine/entity/domain/restaurant/model"
	"github.com/blackhorseya/godine/pkg/contextx"
	"go.uber.org/zap"
)

type restaurantHandler struct {
	restaurantClient biz.RestaurantServiceClient
}

// NewRestaurantServiceHandler is used to create a new RestaurantServiceHandler.
func NewRestaurantServiceHandler(restaurantClient biz.RestaurantServiceClient) bizconnect.RestaurantServiceHandler {
	return &restaurantHandler{
		restaurantClient: restaurantClient,
	}
}

func (i *restaurantHandler) CreateRestaurant(
	ctx context.Context,
	c *connect.Request[biz.CreateRestaurantRequest],
) (*connect.Response[model.Restaurant], error) {
	// TODO: 2024/8/30|sean|implement me
	panic("implement me")
}

func (i *restaurantHandler) ListRestaurants(
	ctx context.Context,
	c *connect.Request[biz.ListRestaurantsRequest],
	c2 *connect.ServerStream[model.Restaurant],
) error {
	// TODO: 2024/8/30|sean|implement me
	panic("implement me")
}

func (i *restaurantHandler) GetRestaurant(
	ctx context.Context,
	c *connect.Request[biz.GetRestaurantRequest],
) (*connect.Response[model.Restaurant], error) {
	// TODO: 2024/8/30|sean|implement me
	panic("implement me")
}

func (i *restaurantHandler) ListRestaurantsNonStream(
	c context.Context,
	req *connect.Request[biz.ListRestaurantsRequest],
) (*connect.Response[biz.ListRestaurantsResponse], error) {
	ctx := contextx.Background()

	resp, err := i.restaurantClient.ListRestaurantsNonStream(c, req.Msg)
	if err != nil {
		ctx.Error("Failed to list restaurants non stream", zap.Error(err))
		return nil, err
	}

	return connect.NewResponse(resp), nil
}
