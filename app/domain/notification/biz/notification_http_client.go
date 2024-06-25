package biz

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/url"

	"github.com/blackhorseya/godine/app/infra/configx"
	"github.com/blackhorseya/godine/app/infra/otelx"
	"github.com/blackhorseya/godine/entity/notification/biz"
	"github.com/blackhorseya/godine/entity/notification/model"
	"github.com/blackhorseya/godine/pkg/contextx"
	"github.com/blackhorseya/godine/pkg/responsex"
	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
)

type notificationHTTPClient struct {
	url    string
	client *http.Client
}

// NewNotificationHTTPClient creates a new notification service.
func NewNotificationHTTPClient() biz.INotificationBiz {
	return &notificationHTTPClient{
		url:    configx.C.NotifyRestful.HTTP.URL,
		client: &http.Client{Transport: otelhttp.NewTransport(http.DefaultTransport)},
	}
}

func (i *notificationHTTPClient) CreateNotification(ctx contextx.Contextx, notification *model.Notification) error {
	// todo: 2024/6/26|sean|implement me
	panic("implement me")
}

func (i *notificationHTTPClient) UpdateNotificationStatus(
	ctx contextx.Contextx,
	notificationID string,
	status string,
) error {
	// todo: 2024/6/26|sean|implement me
	panic("implement me")
}

func (i *notificationHTTPClient) GetNotification(
	ctx contextx.Contextx,
	notificationID string,
) (item *model.Notification, err error) {
	ctx, span := otelx.Span(ctx, "biz.notification.http_client.GetNotification")
	defer span.End()

	ep, err := url.ParseRequestURI(i.url + "/api/v1/notifications/" + notificationID)
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
		Data               *model.Notification `json:"data"`
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

func (i *notificationHTTPClient) ListNotificationsByUser(
	ctx contextx.Contextx,
	userID string,
	options biz.ListNotificationsOptions,
) (items []*model.Notification, total int, err error) {
	// todo: 2024/6/26|sean|implement me
	panic("implement me")
}
