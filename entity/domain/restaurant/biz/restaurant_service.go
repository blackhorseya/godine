//go:generate mockgen -destination=./mock_${GOFILE} -package=${GOPACKAGE} -source=${GOFILE}

package biz

import (
	model2 "github.com/blackhorseya/godine/entity/domain/restaurant/model"
	"github.com/blackhorseya/godine/pkg/contextx"
)

// ListRestaurantsOptions defines the options for listing restaurants.
type ListRestaurantsOptions struct {
	// Page is the page number.
	Page int

	// Size is the number of items per page.
	Size int
}

// IRestaurantBiz defines the business operations for restaurant management.
// Deprecated: use RestaurantServiceServer instead.
type IRestaurantBiz interface {
	// CreateRestaurant creates a new restaurant. This method initializes a new
	// restaurant with the provided name and address.
	CreateRestaurant(ctx contextx.Contextx, name, address string) (item *model2.Restaurant, err error)

	// GetRestaurant retrieves the restaurant with the specified ID.
	GetRestaurant(ctx contextx.Contextx, id string) (item *model2.Restaurant, err error)

	// ListRestaurants retrieves a list of restaurants.
	ListRestaurants(
		ctx contextx.Contextx,
		options ListRestaurantsOptions,
	) (items []*model2.Restaurant, total int, err error)

	// UpdateRestaurant updates the details of an existing restaurant.
	UpdateRestaurant(ctx contextx.Contextx, id string, name string, address *model2.Address) error

	// DeleteRestaurant deletes a restaurant by its ID.
	DeleteRestaurant(ctx contextx.Contextx, id string) error

	// ChangeRestaurantStatus changes the operational status of a restaurant.
	ChangeRestaurantStatus(ctx contextx.Contextx, restaurantID string, isOpen bool) error
}
