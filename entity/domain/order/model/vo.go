package model

// Address represents a value object for an order's delivery address.
type Address struct {
	// Street is the street address of the delivery location.
	Street string `json:"street,omitempty" bson:"street"`

	// City is the city where the order is to be delivered.
	City string `json:"city,omitempty" bson:"city"`

	// State is the state where the order is to be delivered.
	State string `json:"state,omitempty" bson:"state"`

	// ZipCode is the postal code of the delivery location.
	ZipCode string `json:"zip_code,omitempty" bson:"zip_code"`
}

func (a Address) String() string {
	return a.Street + ", " + a.City + ", " + a.State + " " + a.ZipCode
}
