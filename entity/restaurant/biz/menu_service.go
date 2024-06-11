//go:generate mockgen -destination=./mock_${GOFILE} -package=${GOPACKAGE} -source=${GOFILE}

package biz

import (
	"github.com/blackhorseya/godine/entity/restaurant/model"
	"github.com/blackhorseya/godine/pkg/contextx"
	"github.com/google/uuid"
)

// IMenuBiz defines the business operations for menu management.
type IMenuBiz interface {
	// AddMenuItem adds a new menu item to a restaurant's menu.
	AddMenuItem(
		ctx contextx.Contextx,
		restaurantID uuid.UUID,
		name, description string,
		price float64,
	) (item *model.MenuItem, err error)

	// GetMenuItems retrieves all menu items for a specific restaurant.
	GetMenuItems(ctx contextx.Contextx, restaurantID uuid.UUID) (items []model.MenuItem, total int, err error)

	// GetMenuItem retrieves a specific menu item by its ID from a restaurant's menu.
	GetMenuItem(ctx contextx.Contextx, restaurantID, menuItemID uuid.UUID) (item *model.MenuItem, err error)

	// UpdateMenuItem updates the details of an existing menu item.
	UpdateMenuItem(
		ctx contextx.Contextx,
		restaurantID, menuItemID uuid.UUID,
		name, description string,
		price float64,
		isAvailable bool,
	) error

	// RemoveMenuItem removes a menu item from a restaurant's menu.
	RemoveMenuItem(ctx contextx.Contextx, restaurantID, menuItemID uuid.UUID) error
}
