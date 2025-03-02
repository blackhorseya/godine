package domain

import (
	"context"

	"github.com/blackhorseya/godine/app/infra/otelx"
	"github.com/blackhorseya/godine/entity/domain/restaurant/biz"
	"github.com/blackhorseya/godine/entity/domain/restaurant/model"
	"github.com/blackhorseya/godine/entity/domain/restaurant/repo"
	"github.com/blackhorseya/godine/pkg/contextx"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
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
	next, span := otelx.Tracer.Start(c, "menu.biz.AddMenuItem")
	defer span.End()

	ctx := contextx.WithContextx(c)

	restaurant, err := i.restaurants.GetByID(next, req.RestaurantId)
	if err != nil {
		ctx.Error("get restaurant by id failed", zap.Error(err), zap.String("restaurant_id", req.RestaurantId))
		return nil, err
	}

	item, err := restaurant.AddMenuItem(req.Name, req.Description, req.Price)
	if err != nil {
		ctx.Error("add menu item failed", zap.Error(err), zap.Any("request", req))
		return nil, err
	}

	err = i.restaurants.Update(next, restaurant)
	if err != nil {
		ctx.Error("update restaurant failed", zap.Error(err), zap.String("restaurant_id", req.RestaurantId))
		return nil, err
	}

	return item, nil
}

func (i *menuService) GetMenuItem(c context.Context, req *biz.GetMenuItemRequest) (*model.MenuItem, error) {
	next, span := otelx.Tracer.Start(c, "menu.biz.GetMenuItem")
	defer span.End()

	ctx := contextx.WithContextx(c)

	restaurant, err := i.restaurants.GetByID(next, req.RestaurantId)
	if err != nil {
		ctx.Error("get restaurant by id failed", zap.Error(err), zap.String("restaurant_id", req.RestaurantId))
		return nil, err
	}

	for _, item := range restaurant.Menu {
		if item.Id == req.MenuItemId {
			return item, nil
		}
	}

	return nil, status.Errorf(
		codes.NotFound,
		"menu item %s in restaurant %s not found",
		req.MenuItemId,
		req.RestaurantId,
	)
}

func (i *menuService) ListMenuItems(req *biz.ListMenuItemsRequest, stream biz.MenuService_ListMenuItemsServer) error {
	c := stream.Context()
	next, span := otelx.Tracer.Start(c, "menu.biz.ListMenuItems")
	defer span.End()

	ctx := contextx.WithContextx(c)

	restaurant, err := i.restaurants.GetByID(next, req.RestaurantId)
	if err != nil {
		ctx.Error("get restaurant by id failed", zap.Error(err), zap.String("restaurant_id", req.RestaurantId))
		return err
	}

	for _, item := range restaurant.Menu {
		if err = stream.Send(item); err != nil {
			ctx.Error("send menu item failed", zap.Error(err), zap.Any("item", item))
			return err
		}
	}

	return nil
}
