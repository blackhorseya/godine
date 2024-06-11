package biz

import (
	"github.com/blackhorseya/godine/entity/restaurant/biz"
	"github.com/blackhorseya/godine/entity/restaurant/model"
	"github.com/blackhorseya/godine/pkg/contextx"
	"github.com/google/uuid"
)

type menuBiz struct {
}

// NewMenuBiz create and return a new menu biz
func NewMenuBiz() biz.IMenuBiz {
	return &menuBiz{}
}

func (i *menuBiz) AddMenuItem(
	ctx contextx.Contextx,
	restaurantID uuid.UUID,
	name, description string,
	price float64,
) (item *model.MenuItem, err error) {
	// todo: 2024/6/11|sean|implement me
	panic("implement me")
}

func (i *menuBiz) GetMenuItems(
	ctx contextx.Contextx,
	restaurantID uuid.UUID,
) (items []model.MenuItem, total int, err error) {
	// todo: 2024/6/11|sean|implement me
	panic("implement me")
}

func (i *menuBiz) GetMenuItem(
	ctx contextx.Contextx,
	restaurantID, menuItemID uuid.UUID,
) (item *model.MenuItem, err error) {
	// todo: 2024/6/11|sean|implement me
	panic("implement me")
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
