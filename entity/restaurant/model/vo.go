package model

// Address represents a value object for restaurant's address.
type Address struct {
	// Street is the street address of the restaurant.
	Street string `json:"street,omitempty" bson:"street"`

	// City is the city where the restaurant is located.
	City string `json:"city,omitempty" bson:"city"`

	// State is the state where the restaurant is located.
	State string `json:"state,omitempty" bson:"state"`

	// ZipCode is the postal code of the restaurant's location.
	ZipCode string `json:"zipCode,omitempty" bson:"zipCode"`
}

func (a Address) String() string {
	return a.Street + ", " + a.City + ", " + a.State + " " + a.ZipCode
}
