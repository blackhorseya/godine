package model

// OrderEvent represents an order event.
type OrderEvent struct {
	// Name is the name of the event.
	Name string `json:"name,omitempty" bson:"name"`

	// Handler is the handler of the event.
	Handler string `json:"handler,omitempty" bson:"handler"`
}
