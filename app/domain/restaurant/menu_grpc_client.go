package restaurant

import (
	"github.com/blackhorseya/godine/app/infra/transports/grpcx"
	"github.com/blackhorseya/godine/entity/domain/restaurant/biz"
)

// NewMenuServiceClient is used to create a new menu service client.
func NewMenuServiceClient(client *grpcx.Client) (biz.MenuServiceClient, error) {
	conn, err := client.Dial("platform")
	if err != nil {
		return nil, err
	}

	return biz.NewMenuServiceClient(conn), nil
}
