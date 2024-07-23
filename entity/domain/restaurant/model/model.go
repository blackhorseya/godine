package model

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Restaurant represents a restaurant entity.
type Restaurant struct {
	// ID is the unique identifier of the restaurant.
	ID string `json:"id,omitempty" bson:"_id,omitempty"`

	// OwnerID is the unique identifier of the restaurant owner.
	OwnerID string `json:"owner_id,omitempty" bson:"owner_id,omitempty"`

	// Name is the name of the restaurant.
	Name string `json:"name,omitempty" bson:"name"`

	// Address is the address of the restaurant.
	Address Address `json:"address,omitempty" bson:"address"`

	// Menu is the list of menu items available in the restaurant.
	Menu []MenuItem `json:"menu,omitempty" bson:"menu"`

	// IsOpen indicates whether the restaurant is open for business.
	IsOpen bool `json:"is_open" bson:"isOpen"`
}

// NewRestaurant creates a new RestaurantAggregate.
func NewRestaurant(name string, address Address) *Restaurant {
	return &Restaurant{
		ID:      "",
		Name:    name,
		Address: address,
		Menu:    []MenuItem{},
		IsOpen:  false,
	}
}

func (x *Restaurant) UnmarshalBSON(bytes []byte) error {
	type Alias Restaurant
	alias := &struct {
		ID     primitive.ObjectID `bson:"_id"`
		*Alias `bson:",inline"`
	}{
		Alias: (*Alias)(x),
	}

	if err := bson.Unmarshal(bytes, alias); err != nil {
		return err
	}

	x.ID = alias.ID.Hex()

	return nil
}

func (x *Restaurant) MarshalBSON() ([]byte, error) {
	type Alias Restaurant
	alias := &struct {
		ID     primitive.ObjectID `bson:"_id"`
		*Alias `bson:",inline"`
	}{
		Alias: (*Alias)(x),
	}

	id, err := primitive.ObjectIDFromHex(x.ID)
	if err != nil {
		return nil, err
	}
	alias.ID = id

	return bson.Marshal(alias)
}

// AddMenuItem adds a new menu item to the restaurant's menu.
func (x *Restaurant) AddMenuItem(name, description string, price float64) {
	menuItem := MenuItem{
		ID:          primitive.NewObjectID().Hex(),
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
	ID string `json:"id,omitempty" bson:"_id,omitempty"`

	// Name is the name of the menu item.
	Name string `json:"name,omitempty" bson:"name"`

	// Description provides details about the menu item.
	Description string `json:"description,omitempty" bson:"description"`

	// Price is the cost of the menu item.
	Price float64 `json:"price,omitempty" bson:"price"`

	// IsAvailable indicates whether the menu item is available.
	IsAvailable bool `json:"is_available,omitempty" bson:"isAvailable"`
}

func (x *MenuItem) UnmarshalBSON(bytes []byte) error {
	type Alias MenuItem
	alias := &struct {
		ID     primitive.ObjectID `bson:"_id"`
		*Alias `bson:",inline"`
	}{
		Alias: (*Alias)(x),
	}

	if err := bson.Unmarshal(bytes, alias); err != nil {
		return err
	}

	x.ID = alias.ID.Hex()

	return nil
}

func (x *MenuItem) MarshalBSON() ([]byte, error) {
	type Alias MenuItem
	alias := &struct {
		ID     primitive.ObjectID `bson:"_id"`
		*Alias `bson:",inline"`
	}{
		Alias: (*Alias)(x),
	}

	id, err := primitive.ObjectIDFromHex(x.ID)
	if err != nil {
		return nil, err
	}
	alias.ID = id

	return bson.Marshal(alias)
}
