package order

import (
	"fmt"

	"github.com/blackhorseya/godine/app/infra/transports/grpcx"
	"github.com/blackhorseya/godine/entity/domain/order/biz"
)

// NewOrderServiceClient returns a new OrderServiceClient instance.
func NewOrderServiceClient(client *grpcx.Client) (biz.OrderServiceClient, error) {
	conn, err := client.Dial("platform")
	if err != nil {
		return nil, fmt.Errorf("failed to dial platform: %w", err)
	}

	return biz.NewOrderServiceClient(conn), nil
}
