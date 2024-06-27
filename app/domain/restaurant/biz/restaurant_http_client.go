package biz

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/url"
	"strconv"

	"github.com/blackhorseya/godine/adapter/restaurant/restful/v1/restaurants"
	"github.com/blackhorseya/godine/app/infra/configx"
	"github.com/blackhorseya/godine/app/infra/otelx"
	"github.com/blackhorseya/godine/entity/domain/restaurant/biz"
	"github.com/blackhorseya/godine/entity/domain/restaurant/model"
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
) (item *model.Restaurant, err error) {
	ctx, span := otelx.Span(ctx, "biz.restaurant.http_client.CreateRestaurant")
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
		Data               *model.Restaurant `json:"data"`
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

func (i *restaurantHTTPClient) GetRestaurant(ctx contextx.Contextx, id string) (item *model.Restaurant, err error) {
	ctx, span := otelx.Span(ctx, "biz.restaurant.http_client.GetRestaurant")
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
		Data               *model.Restaurant `json:"data"`
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
) (items []*model.Restaurant, total int, err error) {
	ctx, span := otelx.Span(ctx, "biz.restaurant.http_client.ListRestaurants")
	defer span.End()

	ep, err := url.ParseRequestURI(i.url + "/api/v1/restaurants")
	if err != nil {
		return nil, 0, err
	}

	q := ep.Query()
	q.Set("page", strconv.Itoa(options.Page))
	q.Set("limit", strconv.Itoa(options.Size))
	ep.RawQuery = q.Encode()

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, ep.String(), nil)
	if err != nil {
		return nil, 0, err
	}

	resp, err := i.client.Do(req)
	if err != nil {
		return nil, 0, err
	}
	defer resp.Body.Close()

	type response struct {
		responsex.Response `json:",inline"`
		Data               []*model.Restaurant `json:"data"`
	}
	var got response
	err = json.NewDecoder(resp.Body).Decode(&got)
	if err != nil {
		return nil, 0, err
	}

	if got.Code != http.StatusOK {
		return nil, 0, errors.New(got.Message)
	}

	count, err := strconv.Atoi(resp.Header.Get("X-Total-Count"))
	if err != nil {
		return nil, 0, err
	}

	return got.Data, count, nil
}

func (i *restaurantHTTPClient) UpdateRestaurant(
	ctx contextx.Contextx,
	id string,
	name string,
	address model.Address,
) error {
	ctx, span := otelx.Span(ctx, "biz.restaurant.http_client.UpdateRestaurant")
	defer span.End()

	ep, err := url.ParseRequestURI(i.url + "/api/v1/restaurants/" + id)
	if err != nil {
		return err
	}

	payload, err := json.Marshal(model.Restaurant{
		Name:    name,
		Address: address,
	})
	if err != nil {
		return err
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPut, ep.String(), bytes.NewReader(payload))
	if err != nil {
		return err
	}

	resp, err := i.client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	type response struct {
		responsex.Response `json:",inline"`
	}
	var got response
	err = json.NewDecoder(resp.Body).Decode(&got)
	if err != nil {
		return err
	}

	if got.Code != http.StatusOK {
		return errors.New(got.Message)
	}

	return nil
}

func (i *restaurantHTTPClient) DeleteRestaurant(ctx contextx.Contextx, id string) error {
	ctx, span := otelx.Span(ctx, "biz.restaurant.http_client.DeleteRestaurant")
	defer span.End()

	ep, err := url.ParseRequestURI(i.url + "/api/v1/restaurants/" + id)
	if err != nil {
		return err
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodDelete, ep.String(), nil)
	if err != nil {
		return err
	}

	resp, err := i.client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	type response struct {
		responsex.Response `json:",inline"`
	}
	var got response
	err = json.NewDecoder(resp.Body).Decode(&got)
	if err != nil {
		return err
	}

	if got.Code != http.StatusOK {
		return errors.New(got.Message)
	}

	return nil
}

func (i *restaurantHTTPClient) ChangeRestaurantStatus(
	ctx contextx.Contextx,
	restaurantID string,
	isOpen bool,
) error {
	// todo: 2024/6/13|sean|implement me
	panic("implement me")
}
