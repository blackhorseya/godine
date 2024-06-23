package biz

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/url"
	"strconv"

	"github.com/blackhorseya/godine/app/infra/configx"
	"github.com/blackhorseya/godine/app/infra/otelx"
	"github.com/blackhorseya/godine/entity/restaurant/biz"
	"github.com/blackhorseya/godine/entity/restaurant/model"
	"github.com/blackhorseya/godine/pkg/contextx"
	"github.com/blackhorseya/godine/pkg/responsex"
	"github.com/google/uuid"
	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
	"go.uber.org/zap"
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
	ctx, span := otelx.Span(ctx, "restaurant.menuHTTPClient.ListMenuItems")
	defer span.End()

	ep, err := url.ParseRequestURI(i.url + "/api/v1/restaurants/" + restaurantID.String() + "/items")
	if err != nil {
		ctx.Error("parse request uri failed", zap.Error(err))
		return nil, 0, err
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, ep.String(), nil)
	if err != nil {
		ctx.Error("new request failed", zap.Error(err))
		return nil, 0, err
	}

	resp, err := i.client.Do(req)
	if err != nil {
		ctx.Error("do request failed", zap.Error(err))
		return nil, 0, err
	}
	defer resp.Body.Close()

	type response struct {
		responsex.Response `json:",inline"`
		Data               []model.MenuItem `json:"data"`
	}
	var got response
	err = json.NewDecoder(resp.Body).Decode(&got)
	if err != nil {
		ctx.Error("decode response failed", zap.Error(err))
		return nil, 0, err
	}

	if got.Code != http.StatusOK {
		return nil, 0, errors.New(got.Message)
	}

	total, err = strconv.Atoi(resp.Header.Get("X-Total-Count"))
	if err != nil {
		ctx.Error("get total count failed", zap.Error(err))
		return nil, 0, err
	}

	return got.Data, total, nil
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
