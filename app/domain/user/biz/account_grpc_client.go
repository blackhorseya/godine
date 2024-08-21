package biz

import (
	"github.com/blackhorseya/godine/app/infra/transports/grpcx"
	"github.com/blackhorseya/godine/entity/domain/user/biz"
)

// NewAccountServiceClient will create a new account service client.
func NewAccountServiceClient(client *grpcx.Client) (biz.AccountServiceClient, error) {
	conn, err := client.Dial("account")
	if err != nil {
		return nil, err
	}

	return biz.NewAccountServiceClient(conn), nil
}