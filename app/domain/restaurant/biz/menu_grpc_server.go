package biz

import (
	"context"

	"github.com/blackhorseya/godine/entity/domain/restaurant/biz"
	"github.com/blackhorseya/godine/entity/domain/restaurant/model"
)

type menuService struct {
}

// NewMenuService is used to create a new menu service.
func NewMenuService() biz.MenuServiceServer {
	return &menuService{}
}

func (i *menuService) AddMenuItem(ctx context.Context, request *biz.AddMenuItemRequest) (*model.MenuItem, error) {
	// TODO: 2024/8/21|sean|implement me
	panic("implement me")
}

func (i *menuService) GetMenuItem(ctx context.Context, request *biz.GetMenuItemRequest) (*model.MenuItem, error) {
	// TODO: 2024/8/21|sean|implement me
	panic("implement me")
}

func (i *menuService) ListMenuItems(req *biz.ListMenuItemsRequest, stream biz.MenuService_ListMenuItemsServer) error {
	// TODO: 2024/8/21|sean|implement me
	panic("implement me")
}
