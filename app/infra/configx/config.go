package configx

// Configuration defines the configuration struct.
type Configuration struct {
	RestaurantRestful Application `json:"restaurant_restful" yaml:"restaurantRestful" mapstructure:"restaurantRestful"`
}
