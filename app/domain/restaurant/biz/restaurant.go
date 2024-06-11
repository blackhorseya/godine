package biz

import (
	"github.com/blackhorseya/godine/entity/restaurant/biz"
	"github.com/blackhorseya/godine/entity/restaurant/model"
	"github.com/blackhorseya/godine/pkg/contextx"
	"github.com/google/uuid"
)

type restaurantBiz struct {
}

// NewRestaurantBiz create and return a new restaurant biz
func NewRestaurantBiz() biz.IRestaurantBiz {
	return &restaurantBiz{}
}

func (i *restaurantBiz) CreateRestaurant(
	ctx contextx.Contextx,
	name, address string,
) (item *model.Restaurant, err error) {
	// todo: 2024/6/11|sean|implement me
	panic("implement me")
}

func (i *restaurantBiz) GetRestaurant(ctx contextx.Contextx, id uuid.UUID) (item *model.Restaurant, err error) {
	// todo: 2024/6/11|sean|implement me
	panic("implement me")
}

func (i *restaurantBiz) ListRestaurants(
	ctx contextx.Contextx,
	options biz.ListRestaurantsOptions,
) (items []model.Restaurant, total int, err error) {
	// todo: 2024/6/11|sean|implement me
	panic("implement me")
}

func (i *restaurantBiz) UpdateRestaurant(
	ctx contextx.Contextx,
	id uuid.UUID,
	name string,
	address model.Address,
) error {
	// todo: 2024/6/11|sean|implement me
	panic("implement me")
}

func (i *restaurantBiz) DeleteRestaurant(ctx contextx.Contextx, id uuid.UUID) error {
	// todo: 2024/6/11|sean|implement me
	panic("implement me")
}

func (i *restaurantBiz) SearchRestaurants(
	ctx contextx.Contextx,
	keyword string,
) (items []model.Restaurant, total int, err error) {
	// todo: 2024/6/11|sean|implement me
	panic("implement me")
}

func (i *restaurantBiz) ChangeRestaurantStatus(ctx contextx.Contextx, restaurantID uuid.UUID, isOpen bool) error {
	// todo: 2024/6/11|sean|implement me
	panic("implement me")
}
