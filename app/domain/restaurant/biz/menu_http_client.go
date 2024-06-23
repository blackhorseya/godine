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

type menuHTTPClient struct {
	url    string
	client *http.Client
}

// NewMenuHTTPClient is used to create a new menu biz client.
func NewMenuHTTPClient() biz.IMenuBiz {
	return &menuHTTPClient{
		url:    configx.C.RestaurantRestful.HTTP.URL,
		client: &http.Client{Transport: otelhttp.NewTransport(http.DefaultTransport)},
	}
}

func (i *menuHTTPClient) AddMenuItem(
	ctx contextx.Contextx,
	restaurantID uuid.UUID,
	name, description string,
	price float64,
) (item *model.MenuItem, err error) {
	// todo: 2024/6/23|sean|implement me
	panic("implement me")
}

func (i *menuHTTPClient) ListMenuItems(
	ctx contextx.Contextx,
	restaurantID uuid.UUID,
) (items []model.MenuItem, total int, err error) {
	// todo: 2024/6/23|sean|implement me
	panic("implement me")
}

func (i *menuHTTPClient) GetMenuItem(
	ctx contextx.Contextx,
	restaurantID, menuItemID uuid.UUID,
) (item *model.MenuItem, err error) {
	// todo: 2024/6/23|sean|implement me
	panic("implement me")
}

func (i *menuHTTPClient) UpdateMenuItem(
	ctx contextx.Contextx,
	restaurantID, menuItemID uuid.UUID,
	name, description string,
	price float64,
	isAvailable bool,
) error {
	// todo: 2024/6/23|sean|implement me
	panic("implement me")
}

func (i *menuHTTPClient) RemoveMenuItem(
	ctx contextx.Contextx,
	restaurantID, menuItemID uuid.UUID,
) error {
	// todo: 2024/6/23|sean|implement me
	panic("implement me")
}
