package model

// Address represents a delivery address.
type Address struct {
	// Street is the street address.
	Street string `json:"street,omitempty" bson:"street"`

	// City is the city of the address.
	City string `json:"city,omitempty" bson:"city"`

	// State is the state of the address.
	State string `json:"state,omitempty" bson:"state"`

	// ZipCode is the postal code of the address.
	ZipCode string `json:"zip_code,omitempty" bson:"zip_code"`

	// Country is the country of the address.
	Country string `json:"country,omitempty" bson:"country"`
}
