package configx

// Configuration defines the configuration struct.
type Configuration struct {
	RestaurantRestful Application `json:"restaurant_restful" yaml:"restaurantRestful"`
	OrderRestful      Application `json:"order_restful" yaml:"orderRestful"`
	OrderHandler      Application `json:"order_handler" yaml:"orderHandler"`
	UserRestful       Application `json:"user_restful" yaml:"userRestful"`
	LogisticsRestful  Application `json:"logistics_restful" yaml:"logisticsRestful"`
	NotifyRestful     Application `json:"notify_restful" yaml:"notifyRestful"`
}
