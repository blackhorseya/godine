//go:generate mockgen -destination=./mock_${GOFILE} -package=${GOPACKAGE} -source=${GOFILE}

package biz

import (
	"github.com/blackhorseya/godine/entity/restaurant/model"
	"github.com/blackhorseya/godine/pkg/contextx"
)

// ListRestaurantsOptions defines the options for listing restaurants.
type ListRestaurantsOptions struct {
	// Page is the page number.
	Page int

	// PageSize is the number of items per page.
	PageSize int
}

// IRestaurantBiz defines the business operations for restaurant management.
type IRestaurantBiz interface {
	// CreateRestaurant creates a new restaurant. This method initializes a new
	// restaurant with the provided name and address.
	CreateRestaurant(ctx contextx.Contextx, name, address string) (item *model.Restaurant, err error)

	// GetRestaurant retrieves the restaurant with the specified ID.
	GetRestaurant(ctx contextx.Contextx, id string) (item *model.Restaurant, err error)

	// ListRestaurants retrieves a list of restaurants.
	ListRestaurants(
		ctx contextx.Contextx,
		options ListRestaurantsOptions,
	) (items []*model.Restaurant, total int, err error)

	// UpdateRestaurant updates the details of an existing restaurant.
	UpdateRestaurant(ctx contextx.Contextx, id string, name string, address model.Address) error

	// DeleteRestaurant deletes a restaurant by its ID.
	DeleteRestaurant(ctx contextx.Contextx, id string) error

	// ChangeRestaurantStatus changes the operational status of a restaurant.
	ChangeRestaurantStatus(ctx contextx.Contextx, restaurantID string, isOpen bool) error
}
