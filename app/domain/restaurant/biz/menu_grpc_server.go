package biz

import (
	"context"
	"fmt"

	"github.com/blackhorseya/godine/app/infra/otelx"
	"github.com/blackhorseya/godine/entity/domain/restaurant/biz"
	"github.com/blackhorseya/godine/entity/domain/restaurant/model"
	"github.com/blackhorseya/godine/entity/domain/restaurant/repo"
	"github.com/blackhorseya/godine/pkg/contextx"
	"go.uber.org/zap"
)

type menuService struct {
	restaurants repo.IRestaurantRepo
}

// NewMenuService is used to create a new menu service.
func NewMenuService(restaurants repo.IRestaurantRepo) biz.MenuServiceServer {
	return &menuService{
		restaurants: restaurants,
	}
}

func (i *menuService) AddMenuItem(c context.Context, req *biz.AddMenuItemRequest) (*model.MenuItem, error) {
	ctx, err := contextx.FromContext(c)
	if err != nil {
		return nil, fmt.Errorf("failed to get contextx: %w", err)
	}

	ctx, span := otelx.Span(ctx, "menu.biz.AddMenuItem")
	defer span.End()

	restaurant, err := i.restaurants.GetByID(ctx, req.RestaurantId)
	if err != nil {
		ctx.Error("get restaurant by id failed", zap.Error(err), zap.String("restaurant_id", req.RestaurantId))
		return nil, err
	}

	item, err := restaurant.AddMenuItem(req.Name, req.Description, req.Price)
	if err != nil {
		ctx.Error("add menu item failed", zap.Error(err), zap.Any("request", req))
		return nil, err
	}

	err = i.restaurants.Update(ctx, restaurant)
	if err != nil {
		ctx.Error("update restaurant failed", zap.Error(err), zap.String("restaurant_id", req.RestaurantId))
		return nil, err
	}

	return item, nil
}

func (i *menuService) GetMenuItem(ctx context.Context, request *biz.GetMenuItemRequest) (*model.MenuItem, error) {
	// TODO: 2024/8/21|sean|implement me
	panic("implement me")
}

func (i *menuService) ListMenuItems(req *biz.ListMenuItemsRequest, stream biz.MenuService_ListMenuItemsServer) error {
	// TODO: 2024/8/21|sean|implement me
	panic("implement me")
}
