package biz

import (
	"net/http"

	"github.com/blackhorseya/godine/app/infra/configx"
	"github.com/blackhorseya/godine/entity/logistics/biz"
	"github.com/blackhorseya/godine/entity/logistics/model"
	"github.com/blackhorseya/godine/pkg/contextx"
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
	// todo: 2024/6/25|sean|implement me
	panic("implement me")
}

func (i *logisticsHTTPClient) UpdateDeliveryStatus(ctx contextx.Contextx, deliveryID string, status string) error {
	// todo: 2024/6/25|sean|implement me
	panic("implement me")
}

func (i *logisticsHTTPClient) GetDelivery(ctx contextx.Contextx, deliveryID string) (item *model.Delivery, err error) {
	// todo: 2024/6/25|sean|implement me
	panic("implement me")
}

func (i *logisticsHTTPClient) ListDeliveriesByDriver(
	ctx contextx.Contextx,
	driverID string,
	options biz.ListDeliveriesOptions,
) (items []*model.Delivery, total int, err error) {
	// todo: 2024/6/25|sean|implement me
	panic("implement me")
}
