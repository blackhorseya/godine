package configx

// Configuration defines the configuration struct.
type Configuration struct {
	RestaurantRestful Application `json:"restaurant_restful" yaml:"restaurantRestful"`
	OrderRestful      Application `json:"order_restful" yaml:"orderRestful"`
	UserRestful       Application `json:"user_restful" yaml:"userRestful"`
}
