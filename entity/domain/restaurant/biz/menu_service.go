//go:generate mockgen -destination=./mock_${GOFILE} -package=${GOPACKAGE} -source=${GOFILE}

package biz

import (
	"github.com/blackhorseya/godine/entity/domain/restaurant/model"
	"github.com/blackhorseya/godine/pkg/contextx"
)

// IMenuBiz defines the business operations for menu management.
type IMenuBiz interface {
	// AddMenuItem adds a new menu item to a restaurant's menu.
	AddMenuItem(
		ctx contextx.Contextx,
		restaurantID string,
		name, description string,
		price float64,
	) (item *model.MenuItem, err error)

	// ListMenuItems retrieves a list of menu items from a restaurant's menu.
	ListMenuItems(ctx contextx.Contextx, restaurantID string) (items []model.MenuItem, total int, err error)

	// GetMenuItem retrieves a specific menu item by its ID from a restaurant's menu.
	GetMenuItem(ctx contextx.Contextx, restaurantID, menuItemID string) (item *model.MenuItem, err error)

	// UpdateMenuItem updates the details of an existing menu item.
	UpdateMenuItem(
		ctx contextx.Contextx,
		restaurantID, menuItemID string,
		name, description string,
		price float64,
		isAvailable bool,
	) error

	// RemoveMenuItem removes a menu item from a restaurant's menu.
	RemoveMenuItem(ctx contextx.Contextx, restaurantID, menuItemID string) error
}
