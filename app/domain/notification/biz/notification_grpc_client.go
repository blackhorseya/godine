package biz

import (
	"fmt"

	"github.com/blackhorseya/godine/app/infra/transports/grpcx"
	"github.com/blackhorseya/godine/entity/domain/notification/biz"
)

// NewNotificationServiceClient creates a new impl service client.
func NewNotificationServiceClient(client *grpcx.Client) (biz.NotificationServiceClient, error) {
	conn, err := client.Dial("platform")
	if err != nil {
		return nil, fmt.Errorf("failed to dial platform: %w", err)
	}

	return biz.NewNotificationServiceClient(conn), nil
}
