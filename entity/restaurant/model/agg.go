package model

import (
	"github.com/google/uuid"
)

// RestaurantAggregate represents the restaurant aggregate root.
type RestaurantAggregate struct {
	// Restaurant is the core entity of the aggregate.
	Restaurant Restaurant `json:"restaurant" bson:"restaurant"`
}

// NewRestaurant creates a new RestaurantAggregate.
func NewRestaurant(name string, address Address) *RestaurantAggregate {
	return &RestaurantAggregate{
		Restaurant: Restaurant{
			ID:      uuid.New(),
			Name:    name,
			Address: address,
			Menu:    []MenuItem{},
		},
	}
}

// AddMenuItem adds a new menu item to the restaurant's menu.
func (ra *RestaurantAggregate) AddMenuItem(name, description string, price float64) {
	menuItem := MenuItem{
		ID:          uuid.New(),
		Name:        name,
		Description: description,
		Price:       price,
		IsAvailable: true,
	}
	ra.Restaurant.Menu = append(ra.Restaurant.Menu, menuItem)
}
