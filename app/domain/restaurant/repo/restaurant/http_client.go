package restaurant

import (
	"net/http"

	"github.com/blackhorseya/godine/entity/restaurant/model"
	"github.com/blackhorseya/godine/entity/restaurant/repo"
	"github.com/blackhorseya/godine/pkg/contextx"
	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
)

type httpClient struct {
	client *http.Client
}

// NewHTTPClient is used to create a new restaurant repo http client.
func NewHTTPClient() repo.IRestaurantRepo {
	return &httpClient{
		client: &http.Client{Transport: otelhttp.NewTransport(http.DefaultTransport)},
	}
}

func (i *httpClient) Create(ctx contextx.Contextx, data *model.Restaurant) (err error) {
	// todo: 2024/6/13|sean|implement me
	panic("implement me")
}

func (i *httpClient) Update(ctx contextx.Contextx, data *model.Restaurant) (err error) {
	// todo: 2024/6/13|sean|implement me
	panic("implement me")
}

func (i *httpClient) Delete(ctx contextx.Contextx, id string) (err error) {
	// todo: 2024/6/13|sean|implement me
	panic("implement me")
}

func (i *httpClient) GetByID(ctx contextx.Contextx, id string) (item *model.Restaurant, err error) {
	// todo: 2024/6/13|sean|implement me
	panic("implement me")
}

func (i *httpClient) List(
	ctx contextx.Contextx,
	condition repo.ListCondition,
) (items []*model.Restaurant, total int, err error) {
	// todo: 2024/6/13|sean|implement me
	panic("implement me")
}
