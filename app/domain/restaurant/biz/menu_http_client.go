package biz

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/url"
	"strconv"

	"github.com/blackhorseya/godine/app/infra/configx"
	"github.com/blackhorseya/godine/app/infra/otelx"
	"github.com/blackhorseya/godine/entity/domain/restaurant/biz"
	"github.com/blackhorseya/godine/entity/domain/restaurant/model"
	"github.com/blackhorseya/godine/pkg/contextx"
	"github.com/blackhorseya/godine/pkg/errorx"
	"github.com/blackhorseya/godine/pkg/responsex"
	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
	"go.uber.org/zap"
)

const itemRouter = "/items/"

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
	restaurantID string,
	name, description string,
	price float64,
) (item *model.MenuItem, err error) {
	ctx, span := otelx.Span(ctx, "biz.menu.http_client.AddMenuItem")
	defer span.End()

	ep, err := url.ParseRequestURI(i.url + restaurantRouter + restaurantID + "/items")
	if err != nil {
		ctx.Error("parse request uri failed", zap.Error(err))
		return nil, err
	}

	payload, err := json.Marshal(model.MenuItem{
		Name:        name,
		Description: description,
		Price:       price,
	})
	if err != nil {
		ctx.Error("marshal payload failed", zap.Error(err))
		return nil, err
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, ep.String(), bytes.NewReader(payload))
	if err != nil {
		ctx.Error("new request failed", zap.Error(err))
		return nil, err
	}

	resp, err := i.client.Do(req)
	if err != nil {
		ctx.Error("do request failed", zap.Error(err))
		return nil, err
	}
	defer resp.Body.Close()

	type response struct {
		responsex.Response `json:",inline"`
		Data               *model.MenuItem `json:"data"`
	}
	var got response
	err = json.NewDecoder(resp.Body).Decode(&got)
	if err != nil {
		ctx.Error("decode response failed", zap.Error(err))
		return nil, err
	}

	if got.Code != http.StatusOK {
		return nil, errorx.New(got.Code, got.Code, got.Message)
	}

	return got.Data, nil
}

func (i *menuHTTPClient) ListMenuItems(
	ctx contextx.Contextx,
	restaurantID string,
) (items []model.MenuItem, total int, err error) {
	ctx, span := otelx.Span(ctx, "biz.menu.http_client.ListMenuItems")
	defer span.End()

	ep, err := url.ParseRequestURI(i.url + restaurantRouter + restaurantID + "/items")
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
		return nil, 0, errorx.New(got.Code, got.Code, got.Message)
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
	restaurantID, menuItemID string,
) (item *model.MenuItem, err error) {
	ctx, span := otelx.Span(ctx, "biz.menu.http_client.GetMenuItem")
	defer span.End()

	ep, err := url.ParseRequestURI(i.url + restaurantRouter + restaurantID + itemRouter + menuItemID)
	if err != nil {
		ctx.Error("parse request uri failed", zap.Error(err))
		return nil, err
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, ep.String(), nil)
	if err != nil {
		ctx.Error("new request failed", zap.Error(err))
		return nil, err
	}

	resp, err := i.client.Do(req)
	if err != nil {
		ctx.Error("do request failed", zap.Error(err))
		return nil, err
	}
	defer resp.Body.Close()

	type response struct {
		responsex.Response `json:",inline"`
		Data               *model.MenuItem `json:"data"`
	}
	var got response
	err = json.NewDecoder(resp.Body).Decode(&got)
	if err != nil {
		ctx.Error("decode response failed", zap.Error(err))
		return nil, err
	}

	if got.Code != http.StatusOK {
		return nil, errorx.New(got.Code, got.Code, got.Message)
	}

	return got.Data, nil
}

func (i *menuHTTPClient) UpdateMenuItem(
	ctx contextx.Contextx,
	restaurantID, menuItemID string,
	name, description string,
	price float64,
	isAvailable bool,
) error {
	ctx, span := otelx.Span(ctx, "biz.menu.http_client.UpdateMenuItem")
	defer span.End()

	ep, err := url.ParseRequestURI(i.url + restaurantRouter + restaurantID + itemRouter + menuItemID)
	if err != nil {
		ctx.Error("parse request uri failed", zap.Error(err))
		return err
	}

	payload, err := json.Marshal(model.MenuItem{
		Name:        name,
		Description: description,
		Price:       price,
		IsAvailable: isAvailable,
	})
	if err != nil {
		ctx.Error("marshal payload failed", zap.Error(err))
		return err
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPut, ep.String(), bytes.NewReader(payload))
	if err != nil {
		ctx.Error("new request failed", zap.Error(err))
		return err
	}

	resp, err := i.client.Do(req)
	if err != nil {
		ctx.Error("do request failed", zap.Error(err))
		return err
	}
	defer resp.Body.Close()

	type response struct {
		responsex.Response `json:",inline"`
	}
	var got response
	err = json.NewDecoder(resp.Body).Decode(&got)
	if err != nil {
		ctx.Error("decode response failed", zap.Error(err))
		return err
	}

	if got.Code != http.StatusOK {
		return errorx.New(got.Code, got.Code, got.Message)
	}

	return nil
}

func (i *menuHTTPClient) RemoveMenuItem(
	ctx contextx.Contextx,
	restaurantID, menuItemID string,
) error {
	ctx, span := otelx.Span(ctx, "biz.menu.http_client.RemoveMenuItem")
	defer span.End()

	ep, err := url.ParseRequestURI(i.url + restaurantRouter + restaurantID + itemRouter + menuItemID)
	if err != nil {
		ctx.Error("parse request uri failed", zap.Error(err))
		return err
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodDelete, ep.String(), nil)
	if err != nil {
		ctx.Error("new request failed", zap.Error(err))
		return err
	}

	resp, err := i.client.Do(req)
	if err != nil {
		ctx.Error("do request failed", zap.Error(err))
		return err
	}
	defer resp.Body.Close()

	type response struct {
		responsex.Response `json:",inline"`
	}
	var got response
	err = json.NewDecoder(resp.Body).Decode(&got)
	if err != nil {
		ctx.Error("decode response failed", zap.Error(err))
		return err
	}

	if got.Code != http.StatusOK {
		return errorx.New(got.Code, got.Code, got.Message)
	}

	return nil
}
