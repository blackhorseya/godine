package model

// Address represents a value object for user's address.
type Address struct {
	// Street is the street address of the user.
	Street string `json:"street,omitempty" bson:"street"`

	// City is the city where the user is located.
	City string `json:"city,omitempty" bson:"city"`

	// State is the state where the user is located.
	State string `json:"state,omitempty" bson:"state"`

	// ZipCode is the postal code of the user's location.
	ZipCode string `json:"zip_code,omitempty" bson:"zipCode"`
}

func (a Address) String() string {
	return a.Street + ", " + a.City + ", " + a.State + " " + a.ZipCode
}

// RoleType represents the type of role a user has.
type RoleType string
