package domain

import (
	"context"
	"strconv"
	"time"

	"github.com/blackhorseya/godine/app/infra/otelx"
	"github.com/blackhorseya/godine/entity/domain/restaurant/biz"
	"github.com/blackhorseya/godine/entity/domain/restaurant/model"
	"github.com/blackhorseya/godine/entity/domain/restaurant/repo"
	userM "github.com/blackhorseya/godine/entity/domain/user/model"
	"github.com/blackhorseya/godine/pkg/contextx"
	"github.com/blackhorseya/godine/pkg/persistence"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type restaurantService struct {
	restaurants repo.IRestaurantRepo
}

// NewRestaurantService will create a new restaurant service.
func NewRestaurantService(restaurants repo.IRestaurantRepo) biz.RestaurantServiceServer {
	return &restaurantService{
		restaurants: restaurants,
	}
}

func (i *restaurantService) CreateRestaurant(
	c context.Context,
	req *biz.CreateRestaurantRequest,
) (*model.Restaurant, error) {
	next, span := otelx.Tracer.Start(c, "restaurant.biz.CreateRestaurant")
	defer span.End()

	ctx := contextx.WithContextx(c)

	handler, err := userM.FromContext(c)
	if err != nil {
		ctx.Error("failed to get user from context", zap.Error(err))
		return nil, status.Error(codes.Unauthenticated, err.Error())
	}

	restaurant := model.NewRestaurant(req.Name, req.Address)
	restaurant.CreatedBy = handler.Id

	err = i.restaurants.Create(next, restaurant)
	if err != nil {
		ctx.Error("failed to create restaurant", zap.Error(err))
		return nil, status.Error(codes.Internal, err.Error())
	}

	return restaurant, nil
}

func (i *restaurantService) ListRestaurants(
	req *biz.ListRestaurantsRequest,
	stream biz.RestaurantService_ListRestaurantsServer,
) error {
	c := stream.Context()
	next, span := otelx.Tracer.Start(c, "restaurant.biz.ListRestaurants")
	defer span.End()

	ctx := contextx.WithContextx(c)

	items, total, err := i.restaurants.List(next, persistence.Pagination{
		Limit:  req.PageSize,
		Offset: (req.Page - 1) * req.PageSize,
	})
	if err != nil {
		ctx.Error("failed to list restaurants", zap.Error(err))
		return status.Error(codes.Internal, err.Error())
	}

	err = stream.SetHeader(metadata.New(map[string]string{"total": strconv.Itoa(total)}))
	if err != nil {
		ctx.Error("failed to set header", zap.Error(err))
		return status.Error(codes.Internal, err.Error())
	}

	for _, item := range items {
		err = stream.Send(item)
		if err != nil {
			ctx.Error("failed to send restaurant", zap.Error(err))
			return status.Error(codes.Internal, err.Error())
		}
	}

	return nil
}

func (i *restaurantService) GetRestaurant(
	c context.Context,
	req *biz.GetRestaurantRequest,
) (*model.Restaurant, error) {
	next, span := otelx.Tracer.Start(c, "restaurant.biz.GetRestaurant")
	defer span.End()

	return i.restaurants.GetByID(next, req.RestaurantId)
}

func (i *restaurantService) ListRestaurantsNonStream(
	c context.Context,
	req *biz.ListRestaurantsRequest,
) (*biz.ListRestaurantsResponse, error) {
	next, span := otelx.Tracer.Start(c, "restaurant.biz.ListRestaurantsNonStream")
	defer span.End()

	ctx := contextx.WithContextx(c)

	items, total, err := i.restaurants.List(next, persistence.Pagination{
		Limit:  req.PageSize,
		Offset: (req.Page - 1) * req.PageSize,
	})
	if err != nil {
		ctx.Error("failed to list restaurants", zap.Error(err))
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &biz.ListRestaurantsResponse{
		Restaurants: items,
		Total:       int64(total),
	}, nil
}

func (i *restaurantService) PlaceOrder(c context.Context, req *biz.PlaceOrderRequest) (*biz.PlaceOrderResponse, error) {
	next, span := otelx.Tracer.Start(c, "restaurant.biz.PlaceOrder")
	defer span.End()

	ctx := contextx.WithContextx(c)

	restaurant, err := i.restaurants.GetByID(next, req.RestaurantId)
	if restaurant == nil {
		ctx.Error("restaurant not found", zap.Error(err))
		return nil, status.Error(codes.NotFound, err.Error())
	}

	reserved, err := i.reserveInventory(next, restaurant, req.Dishes)
	if !reserved {
		ctx.Error("failed to reserve inventory", zap.Error(err))
		return nil, err
	}

	order := &model.Order{
		Id:           "",
		RestaurantId: restaurant.Id,
		CustomerId:   req.CustomerId,
		Dishes:       req.Dishes,
		Notes:        req.Notes,
		Status:       "reserved",
		EstimatedAt:  timestamppb.New(time.Now().Add(30 * time.Minute)),
		CreatedAt:    nil,
		UpdatedAt:    nil,
	}
	err = i.restaurants.CreateReservation(next, restaurant, order)
	if err != nil {
		ctx.Error("failed to create order", zap.Error(err))
		return nil, err
	}
	ctx.Debug("order created", zap.Any("order", &order))

	return &biz.PlaceOrderResponse{
		OrderId:     order.Id,
		Status:      order.Status,
		EstimatedAt: order.EstimatedAt,
	}, nil
}

func (i *restaurantService) ListOrders(
	req *biz.ListOrdersRequest,
	stream grpc.ServerStreamingServer[model.Order],
) error {
	// TODO: 2024/10/2|sean|implement me
	panic("implement me")
}

func (i *restaurantService) reserveInventory(
	c context.Context,
	restaurant *model.Restaurant,
	dishes []*model.Dish,
) (bool, error) {
	_, span := otelx.Tracer.Start(c, "restaurant.biz.reserveInventory")
	defer span.End()

	ctx := contextx.WithContextx(c)

	for _, menu := range restaurant.Menu {
		for _, dish := range dishes {
			if menu.Id == dish.MenuItemId {
				if menu.Quantity < dish.Quantity {
					ctx.Error("inventory not enough", zap.Any("menu", &menu), zap.Any("dish", &dish))
					return false, status.Error(codes.FailedPrecondition, "inventory not enough")
				}
				menu.Quantity -= dish.Quantity
			}
		}
	}

	return true, nil
}
