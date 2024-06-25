package biz

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/url"

	"github.com/blackhorseya/godine/app/infra/configx"
	"github.com/blackhorseya/godine/app/infra/otelx"
	"github.com/blackhorseya/godine/entity/logistics/biz"
	"github.com/blackhorseya/godine/entity/logistics/model"
	"github.com/blackhorseya/godine/pkg/contextx"
	"github.com/blackhorseya/godine/pkg/responsex"
	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
)

type logisticsHTTPClient struct {
	url    string
	client *http.Client
}

// NewLogisticsHTTPClient will create a new logistics biz
func NewLogisticsHTTPClient() biz.ILogisticsBiz {
	return &logisticsHTTPClient{
		url:    configx.C.LogisticsRestful.HTTP.URL,
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
		return errors.New(got.Message)
	}

	return nil
}

func (i *logisticsHTTPClient) UpdateDeliveryStatus(ctx contextx.Contextx, deliveryID string, status string) error {
	// todo: 2024/6/25|sean|implement me
	panic("implement me")
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
		return nil, errors.New(got.Message)
	}

	return got.Data, nil
}

func (i *logisticsHTTPClient) ListDeliveriesByDriver(
	ctx contextx.Contextx,
	driverID string,
	options biz.ListDeliveriesOptions,
) (items []*model.Delivery, total int, err error) {
	// todo: 2024/6/25|sean|implement me
	panic("implement me")
}
