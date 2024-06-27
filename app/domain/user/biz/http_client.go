package biz

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/url"

	"github.com/blackhorseya/godine/adapter/user/restful/v1/users"
	"github.com/blackhorseya/godine/app/infra/configx"
	"github.com/blackhorseya/godine/app/infra/otelx"
	"github.com/blackhorseya/godine/entity/domain/user/biz"
	"github.com/blackhorseya/godine/entity/domain/user/model"
	"github.com/blackhorseya/godine/pkg/contextx"
	"github.com/blackhorseya/godine/pkg/responsex"
	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
)

const userRouter = "/api/v1/users"

type httpClient struct {
	url    string
	client *http.Client
}

// NewUserHTTPClient is used to create a new user biz http client
func NewUserHTTPClient() biz.IUserBiz {
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
	ctx, span := otelx.Span(ctx, "biz.user.http_client.create_user")
	defer span.End()

	ep, err := url.ParseRequestURI(i.url + userRouter)
	if err != nil {
		return nil, err
	}

	payload, err := json.Marshal(users.PostPayload{
		Name:     name,
		Email:    email,
		Password: password,
		Address:  address,
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
		Data               *model.User `json:"data"`
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

func (i *httpClient) GetUser(ctx contextx.Contextx, id string) (item *model.User, err error) {
	ctx, span := otelx.Span(ctx, "biz.user.http_client.get_user")
	defer span.End()

	ep, err := url.ParseRequestURI(i.url + userRouter + "/" + id)
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
		Data               *model.User `json:"data"`
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
