package biz

import (
	"context"

	restB "github.com/blackhorseya/godine/entity/domain/restaurant/biz"
	restM "github.com/blackhorseya/godine/entity/domain/restaurant/model"
)

type restaurantService struct {
}

// NewRestaurantService will create a new restaurant service.
func NewRestaurantService() restB.RestaurantServiceServer {
	return &restaurantService{}
}

func (i *restaurantService) CreateRestaurant(
	ctx context.Context,
	request *restB.CreateRestaurantRequest,
) (*restM.Restaurant, error) {
	// TODO: 2024/8/21|sean|implement me
	panic("implement me")
}

func (i *restaurantService) GetRestaurant(
	ctx context.Context,
	request *restB.GetRestaurantRequest,
) (*restM.Restaurant, error) {
	// TODO: 2024/8/21|sean|implement me
	panic("implement me")
}

func (i *restaurantService) ListRestaurants(
	request *restB.ListRestaurantsRequest,
	stream restB.RestaurantService_ListRestaurantsServer,
) error {
	// TODO: 2024/8/21|sean|implement me
	panic("implement me")
}
