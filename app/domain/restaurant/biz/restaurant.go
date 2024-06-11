package biz

import (
	"github.com/blackhorseya/godine/app/infra/otelx"
	"github.com/blackhorseya/godine/entity/restaurant/biz"
	"github.com/blackhorseya/godine/entity/restaurant/model"
	"github.com/blackhorseya/godine/entity/restaurant/repo"
	"github.com/blackhorseya/godine/pkg/contextx"
	"github.com/google/uuid"
)

type restaurantBiz struct {
	restaurants repo.IRestaurantRepo
}

// NewRestaurantBiz create and return a new restaurant biz
func NewRestaurantBiz(restaurants repo.IRestaurantRepo) biz.IRestaurantBiz {
	return &restaurantBiz{
		restaurants: restaurants,
	}
}

func (i *restaurantBiz) CreateRestaurant(
	ctx contextx.Contextx,
	name, address string,
) (item *model.Restaurant, err error) {
	ctx, span := otelx.Span(ctx, "biz.restaurant.create_restaurant")
	defer span.End()

	restaurant := model.NewRestaurant(name, model.Address{
		Street: address,
	})

	err = i.restaurants.Create(ctx, restaurant)
	if err != nil {
		return nil, err
	}

	return restaurant, nil
}

func (i *restaurantBiz) GetRestaurant(ctx contextx.Contextx, id uuid.UUID) (item *model.Restaurant, err error) {
	ctx, span := otelx.Span(ctx, "biz.restaurant.get_restaurant")
	defer span.End()

	return i.restaurants.GetByID(ctx, id.String())
}

func (i *restaurantBiz) ListRestaurants(
	ctx contextx.Contextx,
	options biz.ListRestaurantsOptions,
) (items []*model.Restaurant, total int, err error) {
	ctx, span := otelx.Span(ctx, "biz.restaurant.list_restaurants")
	defer span.End()

	return i.restaurants.List(ctx, repo.ListCondition{
		Limit:  options.PageSize,
		Offset: (options.Page - 1) * options.PageSize,
	})
}

func (i *restaurantBiz) UpdateRestaurant(
	ctx contextx.Contextx,
	id uuid.UUID,
	name string,
	address model.Address,
) error {
	ctx, span := otelx.Span(ctx, "biz.restaurant.update_restaurant")
	defer span.End()

	restaurant, err := i.restaurants.GetByID(ctx, id.String())
	if err != nil {
		return err
	}

	restaurant.Name = name
	restaurant.Address = address

	return i.restaurants.Update(ctx, restaurant)
}

func (i *restaurantBiz) DeleteRestaurant(ctx contextx.Contextx, id uuid.UUID) error {
	ctx, span := otelx.Span(ctx, "biz.restaurant.delete_restaurant")
	defer span.End()

	return i.restaurants.Delete(ctx, id.String())
}

func (i *restaurantBiz) ChangeRestaurantStatus(ctx contextx.Contextx, restaurantID uuid.UUID, isOpen bool) error {
	ctx, span := otelx.Span(ctx, "biz.restaurant.change_restaurant_status")
	defer span.End()

	restaurant, err := i.restaurants.GetByID(ctx, restaurantID.String())
	if err != nil {
		return err
	}

	restaurant.IsOpen = isOpen

	return i.restaurants.Update(ctx, restaurant)
}
