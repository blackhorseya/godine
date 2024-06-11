package model

import (
	"github.com/google/uuid"
)

// Restaurant represents a restaurant entity.
type Restaurant struct {
	// ID is the unique identifier of the restaurant.
	ID string `json:"id,omitempty" bson:"_id,omitempty"`

	// Name is the name of the restaurant.
	Name string `json:"name,omitempty" bson:"name"`

	// Address is the address of the restaurant.
	Address Address `json:"address,omitempty" bson:"address"`

	// Menu is the list of menu items available in the restaurant.
	Menu []MenuItem `json:"menu,omitempty" bson:"menu"`

	// IsOpen indicates whether the restaurant is open for business.
	IsOpen bool `json:"is_open,omitempty" bson:"isOpen"`
}

// NewRestaurant creates a new RestaurantAggregate.
func NewRestaurant(name string, address Address) *Restaurant {
	return &Restaurant{
		ID:      uuid.New().String(),
		Name:    name,
		Address: address,
		Menu:    []MenuItem{},
		IsOpen:  false,
	}
}

// AddMenuItem adds a new menu item to the restaurant's menu.
func (x *Restaurant) AddMenuItem(name, description string, price float64) {
	menuItem := MenuItem{
		ID:          uuid.New(),
		Name:        name,
		Description: description,
		Price:       price,
		IsAvailable: true,
	}
	x.Menu = append(x.Menu, menuItem)
}

// MenuItem represents an item in the restaurant's menu.
type MenuItem struct {
	// ID is the unique identifier of the menu item.
	ID uuid.UUID `json:"id,omitempty" bson:"_id,omitempty"`

	// Name is the name of the menu item.
	Name string `json:"name,omitempty" bson:"name"`

	// Description provides details about the menu item.
	Description string `json:"description,omitempty" bson:"description"`

	// Price is the cost of the menu item.
	Price float64 `json:"price,omitempty" bson:"price"`

	// IsAvailable indicates whether the menu item is available.
	IsAvailable bool `json:"is_available,omitempty" bson:"isAvailable"`
}
