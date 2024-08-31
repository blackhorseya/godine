package payment

import (
	"fmt"

	"github.com/blackhorseya/godine/app/infra/transports/grpcx"
	"github.com/blackhorseya/godine/entity/domain/payment/biz"
)

// NewPaymentServiceClient creates a new payment service client.
func NewPaymentServiceClient(client *grpcx.Client) (biz.PaymentServiceClient, error) {
	conn, err := client.Dial("platform")
	if err != nil {
		return nil, fmt.Errorf("failed to dial platform: %w", err)
	}

	return biz.NewPaymentServiceClient(conn), nil
}
