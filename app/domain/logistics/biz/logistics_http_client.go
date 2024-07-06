package biz

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/url"
	"strconv"

	"github.com/blackhorseya/godine/app/infra/configx"
	"github.com/blackhorseya/godine/app/infra/otelx"
	"github.com/blackhorseya/godine/entity/domain/logistics/biz"
	"github.com/blackhorseya/godine/entity/domain/logistics/model"
	"github.com/blackhorseya/godine/pkg/contextx"
	"github.com/blackhorseya/godine/pkg/errorx"
	"github.com/blackhorseya/godine/pkg/responsex"
	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
)

type logisticsHTTPClient struct {
	url    string
	client *http.Client
}

// NewLogisticsHTTPClient will create a new logistics biz
func NewLogisticsHTTPClient(config *configx.Configuration) biz.ILogisticsBiz {
	return &logisticsHTTPClient{
		url:    config.LogisticsRestful.HTTP.URL,
		client: &http.Client{Transport: otelhttp.NewTransport(http.DefaultTransport)},
	}
}

func (i *logisticsHTTPClient) CreateDelivery(ctx contextx.Contextx, delivery *model.Delivery) error {
	ctx, span := otelx.Span(ctx, "biz.logistics.http_client.CreateDelivery")
	defer span.End()

	ep, err := url.ParseRequestURI(i.url + "/api/v1/deliveries")
	if err != nil {
		return err
	}

	payload, err := json.Marshal(delivery)
	if err != nil {
		return err
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, ep.String(), bytes.NewReader(payload))
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
		Data               *model.Delivery `json:"data"`
	}
	var got response
	err = json.NewDecoder(resp.Body).Decode(&got)
	if err != nil {
		return err
	}

	if got.Code != http.StatusOK {
		return errorx.New(got.Code, got.Code, got.Message)
	}

	delivery.ID = got.Data.ID

	return nil
}

func (i *logisticsHTTPClient) UpdateDeliveryStatus(ctx contextx.Contextx, deliveryID string, status string) error {
	ctx, span := otelx.Span(ctx, "biz.logistics.http_client.UpdateDeliveryStatus")
	defer span.End()

	ep, err := url.ParseRequestURI(i.url + "/api/v1/deliveries/" + deliveryID + "/status")
	if err != nil {
		return err
	}

	payload, err := json.Marshal(map[string]string{"status": status})
	if err != nil {
		return err
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPatch, ep.String(), bytes.NewReader(payload))
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
		Data               *model.Delivery `json:"data"`
	}
	var got response
	err = json.NewDecoder(resp.Body).Decode(&got)
	if err != nil {
		return err
	}

	if got.Code != http.StatusOK {
		return errorx.New(got.Code, got.Code, got.Message)
	}

	return nil
}

func (i *logisticsHTTPClient) GetDelivery(ctx contextx.Contextx, deliveryID string) (item *model.Delivery, err error) {
	ctx, span := otelx.Span(ctx, "biz.logistics.http_client.GetDelivery")
	defer span.End()

	ep, err := url.ParseRequestURI(i.url + "/api/v1/deliveries/" + deliveryID)
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
		Data               *model.Delivery `json:"data"`
	}
	var got response
	err = json.NewDecoder(resp.Body).Decode(&got)
	if err != nil {
		return nil, err
	}

	if got.Code != http.StatusOK {
		return nil, errorx.New(got.Code, got.Code, got.Message)
	}

	return got.Data, nil
}

func (i *logisticsHTTPClient) ListDeliveriesByDriver(
	ctx contextx.Contextx,
	driverID string,
	options biz.ListDeliveriesOptions,
) (items []*model.Delivery, total int, err error) {
	ctx, span := otelx.Span(ctx, "biz.logistics.http_client.ListDeliveriesByDriver")
	defer span.End()

	ep, err := url.ParseRequestURI(i.url + "/api/v1/deliveries")
	if err != nil {
		return nil, 0, err
	}

	q := ep.Query()
	q.Set("driver_id", driverID)
	q.Set("page", strconv.Itoa(options.Page))
	q.Set("size", strconv.Itoa(options.Size))
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
		Data               []*model.Delivery `json:"data"`
	}
	var got response
	err = json.NewDecoder(resp.Body).Decode(&got)
	if err != nil {
		return nil, 0, err
	}

	if got.Code != http.StatusOK {
		return nil, 0, errorx.New(got.Code, got.Code, got.Message)
	}

	count, err := strconv.Atoi(resp.Header.Get("X-Total-Count"))
	if err != nil {
		return nil, 0, err
	}

	return got.Data, count, nil
}
