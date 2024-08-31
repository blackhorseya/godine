package model

import (
	"errors"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// NewRestaurant creates a new RestaurantAggregate.
func NewRestaurant(name string, address *Address) *Restaurant {
	return &Restaurant{
		Name:    name,
		Address: address,
		IsOpen:  true,
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

	x.Id = alias.ID.Hex()

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

	id, err := primitive.ObjectIDFromHex(x.Id)
	if err != nil {
		return nil, err
	}
	alias.ID = id

	return bson.Marshal(alias)
}

// AddMenuItem adds a new menu item to the restaurant's menu.
func (x *Restaurant) AddMenuItem(name, description string, price float64) (*MenuItem, error) {
	if price <= 0 {
		return nil, errors.New("price must be greater than 0")
	}

	menuItem := &MenuItem{
		Id:          primitive.NewObjectID().Hex(),
		Name:        name,
		Description: description,
		Price:       price,
		IsAvailable: true,
	}
	x.Menu = append(x.Menu, menuItem)

	return menuItem, nil
}
