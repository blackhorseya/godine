package biz

import (
	"fmt"

	"github.com/blackhorseya/godine/app/infra/transports/grpcx"
	"github.com/blackhorseya/godine/entity/domain/logistics/biz"
)

// NewLogisticsServiceClient creates a new logistics service client.
func NewLogisticsServiceClient(client *grpcx.Client) (biz.LogisticsServiceClient, error) {
	conn, err := client.Dial("platform")
	if err != nil {
		return nil, fmt.Errorf("failed to dial platform: %w", err)
	}

	return biz.NewLogisticsServiceClient(conn), nil
}
