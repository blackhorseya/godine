package biz

import (
	"net/http"

	"github.com/blackhorseya/godine/app/infra/otelx"
	"github.com/blackhorseya/godine/entity/restaurant/biz"
	"github.com/blackhorseya/godine/entity/restaurant/model"
	"github.com/blackhorseya/godine/entity/restaurant/repo"
	"github.com/blackhorseya/godine/pkg/contextx"
	"github.com/blackhorseya/godine/pkg/errorx"
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
	restaurantID string,
	name, description string,
	price float64,
) (item *model.MenuItem, err error) {
	ctx, span := otelx.Span(ctx, "menu.biz.add_menu_item")
	defer span.End()

	restaurant, err := i.restaurants.GetByID(ctx, restaurantID)
	if err != nil {
		ctx.Error(
			"get restaurant by id failed",
			zap.Error(err),
			zap.String("restaurant_id", restaurantID),
		)
		return nil, err
	}
	if restaurant == nil {
		ctx.Error("restaurant not found", zap.String("restaurant_id", restaurantID))
		return nil, errorx.New(http.StatusNotFound, 404, "restaurant not found")
	}

	restaurant.AddMenuItem(name, description, price)

	err = i.restaurants.Update(ctx, restaurant)
	if err != nil {
		ctx.Error(
			"update restaurant failed",
			zap.Error(err),
			zap.String("restaurant_id", restaurantID),
		)
		return nil, err
	}

	return &restaurant.Menu[len(restaurant.Menu)-1], nil
}

func (i *menuBiz) ListMenuItems(
	ctx contextx.Contextx,
	restaurantID string,
) (items []model.MenuItem, total int, err error) {
	ctx, span := otelx.Span(ctx, "menu.biz.get_menu_items")
	defer span.End()

	restaurant, err := i.restaurants.GetByID(ctx, restaurantID)
	if err != nil {
		ctx.Error(
			"get restaurant by id failed",
			zap.Error(err),
			zap.String("restaurant_id", restaurantID),
		)
		return nil, 0, err
	}
	if restaurant == nil {
		ctx.Error("restaurant not found", zap.String("restaurant_id", restaurantID))
		return nil, 0, errorx.New(http.StatusNotFound, 404, "restaurant not found")
	}

	return restaurant.Menu, len(restaurant.Menu), nil
}

func (i *menuBiz) GetMenuItem(
	ctx contextx.Contextx,
	restaurantID, menuItemID string,
) (item *model.MenuItem, err error) {
	ctx, span := otelx.Span(ctx, "menu.biz.get_menu_item")
	defer span.End()

	restaurant, err := i.restaurants.GetByID(ctx, restaurantID)
	if err != nil {
		ctx.Error(
			"get restaurant by id failed",
			zap.Error(err),
			zap.String("restaurant_id", restaurantID),
		)
		return nil, err
	}
	if restaurant == nil {
		ctx.Error("restaurant not found", zap.String("restaurant_id", restaurantID))
		return nil, errorx.New(http.StatusNotFound, 404, "restaurant not found")
	}

	for _, menuItem := range restaurant.Menu {
		if menuItem.ID == menuItemID {
			return &menuItem, nil
		}
	}

	return nil, errorx.New(http.StatusNotFound, 404, "menu item not found")
}

func (i *menuBiz) UpdateMenuItem(
	ctx contextx.Contextx,
	restaurantID, menuItemID string,
	name, description string,
	price float64,
	isAvailable bool,
) error {
	ctx, span := otelx.Span(ctx, "menu.biz.update_menu_item")
	defer span.End()

	restaurant, err := i.restaurants.GetByID(ctx, restaurantID)
	if err != nil {
		ctx.Error(
			"get restaurant by id failed",
			zap.Error(err),
			zap.String("restaurant_id", restaurantID),
		)
		return err
	}
	if restaurant == nil {
		ctx.Error("restaurant not found", zap.String("restaurant_id", restaurantID))
		return errorx.New(http.StatusNotFound, 404, "restaurant not found")
	}

	for idx, menuItem := range restaurant.Menu {
		if menuItem.ID == menuItemID {
			restaurant.Menu[idx].Name = name
			restaurant.Menu[idx].Description = description
			restaurant.Menu[idx].Price = price
			restaurant.Menu[idx].IsAvailable = isAvailable

			err = i.restaurants.Update(ctx, restaurant)
			if err != nil {
				ctx.Error(
					"update restaurant failed",
					zap.Error(err),
					zap.String("restaurant_id", restaurantID),
				)
				return err
			}

			return nil
		}
	}

	return errorx.New(http.StatusNotFound, 404, "menu item not found")
}

func (i *menuBiz) RemoveMenuItem(ctx contextx.Contextx, restaurantID, menuItemID string) error {
	ctx, span := otelx.Span(ctx, "menu.biz.remove_menu_item")
	defer span.End()

	restaurant, err := i.restaurants.GetByID(ctx, restaurantID)
	if err != nil {
		ctx.Error(
			"get restaurant by id failed",
			zap.Error(err),
			zap.String("restaurant_id", restaurantID),
		)
		return err
	}
	if restaurant == nil {
		ctx.Error("restaurant not found", zap.String("restaurant_id", restaurantID))
		return errorx.New(http.StatusNotFound, 404, "restaurant not found")
	}

	for idx, menuItem := range restaurant.Menu {
		if menuItem.ID == menuItemID {
			restaurant.Menu = append(restaurant.Menu[:idx], restaurant.Menu[idx+1:]...)

			err = i.restaurants.Update(ctx, restaurant)
			if err != nil {
				ctx.Error(
					"update restaurant failed",
					zap.Error(err),
					zap.String("restaurant_id", restaurantID),
				)
				return err
			}

			return nil
		}
	}

	return errorx.New(http.StatusNotFound, 404, "menu item not found")
}
