package biz

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/url"

	"github.com/blackhorseya/godine/adapter/restaurant/restful/v1/restaurants"
	"github.com/blackhorseya/godine/app/infra/configx"
	"github.com/blackhorseya/godine/app/infra/otelx"
	"github.com/blackhorseya/godine/entity/domain/restaurant/biz"
	model2 "github.com/blackhorseya/godine/entity/domain/restaurant/model"
	"github.com/blackhorseya/godine/pkg/contextx"
	"github.com/blackhorseya/godine/pkg/responsex"
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
) (item *model2.Restaurant, err error) {
	ctx, span := otelx.Span(ctx, "restaurantHTTPClient.CreateRestaurant")
	defer span.End()

	ep, err := url.ParseRequestURI(i.url + "/api/v1/restaurants")
	if err != nil {
		return nil, err
	}

	payload, err := json.Marshal(restaurants.PostPayload{
		Name:        name,
		Description: address,
	})
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, ep.String(), bytes.NewReader(payload))
	if err != nil {
		return nil, err
	}

	resp, err := i.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	type response struct {
		responsex.Response `json:",inline"`
		Data               *model2.Restaurant `json:"data"`
	}
	var got response
	err = json.NewDecoder(resp.Body).Decode(&got)
	if err != nil {
		return nil, err
	}

	if got.Code != http.StatusOK {
		return nil, errors.New(got.Message)
	}

	return got.Data, nil
}

func (i *restaurantHTTPClient) GetRestaurant(ctx contextx.Contextx, id string) (item *model2.Restaurant, err error) {
	ctx, span := otelx.Span(ctx, "restaurantHTTPClient.GetRestaurant")
	defer span.End()

	ep, err := url.ParseRequestURI(i.url + "/api/v1/restaurants/" + id)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, ep.String(), nil)
	if err != nil {
		return nil, err
	}

	resp, err := i.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	type response struct {
		responsex.Response `json:",inline"`
		Data               *model2.Restaurant `json:"data"`
	}
	var got response
	err = json.NewDecoder(resp.Body).Decode(&got)
	if err != nil {
		return nil, err
	}

	if got.Code != http.StatusOK {
		return nil, errors.New(got.Message)
	}

	return got.Data, nil
}

func (i *restaurantHTTPClient) ListRestaurants(
	ctx contextx.Contextx,
	options biz.ListRestaurantsOptions,
) (items []*model2.Restaurant, total int, err error) {
	// todo: 2024/6/13|sean|implement me
	panic("implement me")
}

func (i *restaurantHTTPClient) UpdateRestaurant(
	ctx contextx.Contextx,
	id string,
	name string,
	address model2.Address,
) error {
	// todo: 2024/6/13|sean|implement me
	panic("implement me")
}

func (i *restaurantHTTPClient) DeleteRestaurant(ctx contextx.Contextx, id string) error {
	// todo: 2024/6/13|sean|implement me
	panic("implement me")
}

func (i *restaurantHTTPClient) ChangeRestaurantStatus(
	ctx contextx.Contextx,
	restaurantID string,
	isOpen bool,
) error {
	// todo: 2024/6/13|sean|implement me
	panic("implement me")
}
