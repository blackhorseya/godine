package biz

import (
	"net/http"

	"github.com/blackhorseya/godine/app/infra/configx"
	"github.com/blackhorseya/godine/entity/notification/biz"
	"github.com/blackhorseya/godine/entity/notification/model"
	"github.com/blackhorseya/godine/pkg/contextx"
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
	// todo: 2024/6/26|sean|implement me
	panic("implement me")
}

func (i *notificationHTTPClient) ListNotificationsByUser(
	ctx contextx.Contextx,
	userID string,
	options biz.ListNotificationsOptions,
) (items []*model.Notification, total int, err error) {
	// todo: 2024/6/26|sean|implement me
	panic("implement me")
}
