package biz

import (
	"net/http"

	"github.com/blackhorseya/godine/app/infra/otelx"
	"github.com/blackhorseya/godine/entity/restaurant/biz"
	"github.com/blackhorseya/godine/entity/restaurant/model"
	"github.com/blackhorseya/godine/entity/restaurant/repo"
	"github.com/blackhorseya/godine/pkg/contextx"
	"github.com/blackhorseya/godine/pkg/errorx"
	"github.com/google/uuid"
	"go.uber.org/zap"
)

type menuBiz struct {
	restaurants repo.IRestaurantRepo
}

// NewMenuBiz create and return a new menu biz
func NewMenuBiz(restaurants repo.IRestaurantRepo) biz.IMenuBiz {
	return &menuBiz{restaurants: restaurants}
}

func (i *menuBiz) AddMenuItem(
	ctx contextx.Contextx,
	restaurantID uuid.UUID,
	name, description string,
	price float64,
) (item *model.MenuItem, err error) {
	ctx, span := otelx.Span(ctx, "menu.biz.add_menu_item")
	defer span.End()

	restaurant, err := i.restaurants.GetByID(ctx, restaurantID.String())
	if err != nil {
		ctx.Error(
			"get restaurant by id failed",
			zap.Error(err),
			zap.String("restaurant_id", restaurantID.String()),
		)
		return nil, err
	}
	if restaurant == nil {
		ctx.Error("restaurant not found", zap.String("restaurant_id", restaurantID.String()))
		return nil, errorx.New(http.StatusNotFound, 404, "restaurant not found")
	}

	restaurant.AddMenuItem(name, description, price)

	err = i.restaurants.Update(ctx, restaurant)
	if err != nil {
		ctx.Error(
			"update restaurant failed",
			zap.Error(err),
			zap.String("restaurant_id", restaurantID.String()),
		)
		return nil, err
	}

	return &restaurant.Menu[len(restaurant.Menu)-1], nil
}

func (i *menuBiz) GetMenuItems(
	ctx contextx.Contextx,
	restaurantID uuid.UUID,
) (items []model.MenuItem, total int, err error) {
	ctx, span := otelx.Span(ctx, "menu.biz.get_menu_items")
	defer span.End()

	restaurant, err := i.restaurants.GetByID(ctx, restaurantID.String())
	if err != nil {
		ctx.Error(
			"get restaurant by id failed",
			zap.Error(err),
			zap.String("restaurant_id", restaurantID.String()),
		)
		return nil, 0, err
	}
	if restaurant == nil {
		ctx.Error("restaurant not found", zap.String("restaurant_id", restaurantID.String()))
		return nil, 0, errorx.New(http.StatusNotFound, 404, "restaurant not found")
	}

	return restaurant.Menu, len(restaurant.Menu), nil
}

func (i *menuBiz) GetMenuItem(
	ctx contextx.Contextx,
	restaurantID, menuItemID uuid.UUID,
) (item *model.MenuItem, err error) {
	ctx, span := otelx.Span(ctx, "menu.biz.get_menu_item")
	defer span.End()

	restaurant, err := i.restaurants.GetByID(ctx, restaurantID.String())
	if err != nil {
		ctx.Error(
			"get restaurant by id failed",
			zap.Error(err),
			zap.String("restaurant_id", restaurantID.String()),
		)
		return nil, err
	}
	if restaurant == nil {
		ctx.Error("restaurant not found", zap.String("restaurant_id", restaurantID.String()))
		return nil, errorx.New(http.StatusNotFound, 404, "restaurant not found")
	}

	for _, menuItem := range restaurant.Menu {
		if menuItem.ID == menuItemID.String() {
			return &menuItem, nil
		}
	}

	return nil, errorx.New(http.StatusNotFound, 404, "menu item not found")
}

func (i *menuBiz) UpdateMenuItem(
	ctx contextx.Contextx,
	restaurantID, menuItemID uuid.UUID,
	name, description string,
	price float64,
	isAvailable bool,
) error {
	// todo: 2024/6/11|sean|implement me
	panic("implement me")
}

func (i *menuBiz) RemoveMenuItem(ctx contextx.Contextx, restaurantID, menuItemID uuid.UUID) error {
	// todo: 2024/6/11|sean|implement me
	panic("implement me")
}
