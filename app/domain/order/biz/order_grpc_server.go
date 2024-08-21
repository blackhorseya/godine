package biz

import (
	"context"

	"github.com/blackhorseya/godine/entity/domain/order/biz"
	"github.com/blackhorseya/godine/entity/domain/order/model"
	"github.com/blackhorseya/godine/entity/domain/order/repo"
)

type orderService struct {
	orders repo.IOrderRepo
}

// NewOrderService returns the order service instance.
func NewOrderService(orders repo.IOrderRepo) biz.OrderServiceServer {
	return &orderService{
		orders: orders,
	}
}

func (i *orderService) SubmitOrder(c context.Context, req *biz.SubmitOrderRequest) (*model.Order, error) {
	// TODO: 2024/8/21|sean|implement me
	panic("implement me")
}

func (i *orderService) ListOrders(req *biz.ListOrdersRequest, stream biz.OrderService_ListOrdersServer) error {
	// TODO: 2024/8/21|sean|implement me
	panic("implement me")
}
