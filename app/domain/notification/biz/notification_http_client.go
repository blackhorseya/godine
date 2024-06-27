package biz

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/url"
	"strconv"

	"github.com/blackhorseya/godine/app/infra/configx"
	"github.com/blackhorseya/godine/app/infra/otelx"
	"github.com/blackhorseya/godine/entity/domain/notification/biz"
	"github.com/blackhorseya/godine/entity/domain/notification/model"
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
	ctx, span := otelx.Span(ctx, "biz.notification.http_client.CreateNotification")
	defer span.End()

	ep, err := url.ParseRequestURI(i.url + "/api/v1/notifications")
	if err != nil {
		return err
	}

	payload, err := json.Marshal(notification)
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
		Data               *model.Notification `json:"data"`
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
	ctx, span := otelx.Span(ctx, "biz.notification.http_client.ListNotificationsByUser")
	defer span.End()

	ep, err := url.ParseRequestURI(i.url + "/api/v1/notifications")
	if err != nil {
		return nil, 0, err
	}

	q := ep.Query()
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
		Data               []*model.Notification `json:"data"`
	}
	var got response
	err = json.NewDecoder(resp.Body).Decode(&got)
	if err != nil {
		return nil, 0, err
	}

	if got.Code != http.StatusOK {
		return nil, 0, errors.New(got.Message)
	}

	count, err := strconv.Atoi(resp.Header.Get("X-Total-Count"))
	if err != nil {
		return nil, 0, err
	}

	return got.Data, count, nil
}
