package biz

import (
	"net/http"

	"github.com/blackhorseya/godine/app/infra/configx"
	"github.com/blackhorseya/godine/entity/user/biz"
	"github.com/blackhorseya/godine/entity/user/model"
	"github.com/blackhorseya/godine/pkg/contextx"
	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
)

type httpClient struct {
	url    string
	client *http.Client
}

// NewHTTPClient is used to create a new user biz http client
func NewHTTPClient() biz.IUserBiz {
	return &httpClient{
		url:    configx.C.UserRestful.HTTP.URL,
		client: &http.Client{Transport: otelhttp.NewTransport(http.DefaultTransport)},
	}
}

func (i *httpClient) CreateUser(
	ctx contextx.Contextx,
	name, email, password string,
	address model.Address,
) (item *model.User, err error) {
	// todo: 2024/6/14|sean|implement me
	panic("implement me")
}

func (i *httpClient) GetUser(ctx contextx.Contextx, id string) (item *model.User, err error) {
	// todo: 2024/6/14|sean|implement me
	panic("implement me")
}

func (i *httpClient) ListUsers(
	ctx contextx.Contextx,
	options biz.ListUsersOptions,
) (items []*model.User, total int, err error) {
	// todo: 2024/6/14|sean|implement me
	panic("implement me")
}

func (i *httpClient) UpdateUser(
	ctx contextx.Contextx,
	id string,
	name, email, password string,
	address model.Address,
) error {
	// todo: 2024/6/14|sean|implement me
	panic("implement me")
}

func (i *httpClient) DeleteUser(ctx contextx.Contextx, id string) error {
	// todo: 2024/6/14|sean|implement me
	panic("implement me")
}

func (i *httpClient) ChangeUserStatus(ctx contextx.Contextx, userID string, isActive bool) error {
	// todo: 2024/6/14|sean|implement me
	panic("implement me")
}
