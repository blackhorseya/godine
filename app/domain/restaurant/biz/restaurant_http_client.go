package biz

import (
	"net/http"

	"github.com/blackhorseya/godine/app/infra/configx"
	"github.com/blackhorseya/godine/entity/restaurant/biz"
	"github.com/blackhorseya/godine/entity/restaurant/model"
	"github.com/blackhorseya/godine/pkg/contextx"
	"github.com/google/uuid"
	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
)

type restaurantHTTPClient struct {
	url    string
	client *http.Client
}

// NewRestaurantHTTPClient is used to create a new restaurant biz client.
func NewRestaurantHTTPClient() biz.IRestaurantBiz {
	return &restaurantHTTPClient{
		url:    configx.C.RestaurantRestful.HTTP.URL,
		client: &http.Client{Transport: otelhttp.NewTransport(http.DefaultTransport)},
	}
}

func (i *restaurantHTTPClient) CreateRestaurant(
	ctx contextx.Contextx,
	name, address string,
) (item *model.Restaurant, err error) {
	// todo: 2024/6/13|sean|implement me
	panic("implement me")
}

func (i *restaurantHTTPClient) GetRestaurant(ctx contextx.Contextx, id uuid.UUID) (item *model.Restaurant, err error) {
	// todo: 2024/6/13|sean|implement me
	panic("implement me")
}

func (i *restaurantHTTPClient) ListRestaurants(
	ctx contextx.Contextx,
	options biz.ListRestaurantsOptions,
) (items []*model.Restaurant, total int, err error) {
	// todo: 2024/6/13|sean|implement me
	panic("implement me")
}

func (i *restaurantHTTPClient) UpdateRestaurant(
	ctx contextx.Contextx,
	id uuid.UUID,
	name string,
	address model.Address,
) error {
	// todo: 2024/6/13|sean|implement me
	panic("implement me")
}

func (i *restaurantHTTPClient) DeleteRestaurant(ctx contextx.Contextx, id uuid.UUID) error {
	// todo: 2024/6/13|sean|implement me
	panic("implement me")
}

func (i *restaurantHTTPClient) ChangeRestaurantStatus(
	ctx contextx.Contextx,
	restaurantID uuid.UUID,
	isOpen bool,
) error {
	// todo: 2024/6/13|sean|implement me
	panic("implement me")
}
