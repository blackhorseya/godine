package biz

import (
	"github.com/blackhorseya/godine/app/infra/otelx"
	"github.com/blackhorseya/godine/entity/domain/restaurant/biz"
	"github.com/blackhorseya/godine/entity/domain/restaurant/model"
	"github.com/blackhorseya/godine/entity/domain/restaurant/repo"
	model2 "github.com/blackhorseya/godine/entity/domain/user/model"
	"github.com/blackhorseya/godine/pkg/contextx"
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

	handler, err := model2.FromContext(ctx)
	if err != nil {
		return nil, err
	}

	restaurant := model.NewRestaurant(name, &model.Address{
		Street: address,
	})
	restaurant.CreatedBy = handler.Id

	err = i.restaurants.Create(ctx, restaurant)
	if err != nil {
		return nil, err
	}

	return restaurant, nil
}

func (i *restaurantBiz) GetRestaurant(ctx contextx.Contextx, id string) (item *model.Restaurant, err error) {
	ctx, span := otelx.Span(ctx, "biz.restaurant.get_restaurant")
	defer span.End()

	return i.restaurants.GetByID(ctx, id)
}

func (i *restaurantBiz) ListRestaurants(
	ctx contextx.Contextx,
	options biz.ListRestaurantsOptions,
) (items []*model.Restaurant, total int, err error) {
	ctx, span := otelx.Span(ctx, "biz.restaurant.list_restaurants")
	defer span.End()

	return i.restaurants.List(ctx, repo.ListCondition{
		Limit:  int64(options.Size),
		Offset: int64((options.Page - 1) * options.Size),
	})
}

func (i *restaurantBiz) UpdateRestaurant(
	ctx contextx.Contextx,
	id string,
	name string,
	address *model.Address,
) error {
	ctx, span := otelx.Span(ctx, "biz.restaurant.update_restaurant")
	defer span.End()

	restaurant, err := i.restaurants.GetByID(ctx, id)
	if err != nil {
		return err
	}

	restaurant.Name = name
	restaurant.Address = address

	return i.restaurants.Update(ctx, restaurant)
}

func (i *restaurantBiz) DeleteRestaurant(ctx contextx.Contextx, id string) error {
	ctx, span := otelx.Span(ctx, "biz.restaurant.delete_restaurant")
	defer span.End()

	return i.restaurants.Delete(ctx, id)
}

func (i *restaurantBiz) ChangeRestaurantStatus(ctx contextx.Contextx, restaurantID string, isOpen bool) error {
	ctx, span := otelx.Span(ctx, "biz.restaurant.change_restaurant_status")
	defer span.End()

	restaurant, err := i.restaurants.GetByID(ctx, restaurantID)
	if err != nil {
		return err
	}

	restaurant.IsOpen = isOpen

	return i.restaurants.Update(ctx, restaurant)
}
