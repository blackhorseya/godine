package domain

import (
	"github.com/blackhorseya/godine/app/infra/transports/grpcx"
	restB "github.com/blackhorseya/godine/entity/domain/restaurant/biz"
)

// NewRestaurantServiceClient will create a new restaurant service client.
func NewRestaurantServiceClient(client *grpcx.Client) (restB.RestaurantServiceClient, error) {
	conn, err := client.Dial("platform")
	if err != nil {
		return nil, err
	}

	return restB.NewRestaurantServiceClient(conn), nil
}
