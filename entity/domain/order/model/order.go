package model

import (
	"errors"

	"github.com/blackhorseya/godine/pkg/contextx"
	"google.golang.org/protobuf/types/known/timestamppb"
)

// NewOrder creates a new order.
func NewOrder(userID, restaurantID string, items []*OrderItem) *Order {
	totalAmount := 0.0
	for _, item := range items {
		totalAmount += item.Price * float64(item.Quantity)
	}

	return &Order{
		UserId:       userID,
		RestaurantId: restaurantID,
		Items:        items,
		Status:       OrderStatus_ORDER_STATUS_PENDING,
		CreatedAt:    timestamppb.Now(),
		UpdatedAt:    timestamppb.Now(),
	}
}

// Next transitions the order to the next state.
func (x *Order) Next(ctx contextx.Contextx) (event *OrderEvent, err error) {
	// TODO: 2024/8/21|sean|implement the order state transition logic
	return nil, errors.New("not implemented")
}

// // BeforeSave GORM hook - convert OrderState to string before saving
// func (x *Order) BeforeSave(tx *gorm.DB) (err error) {
// 	if x.ID != "" {
// 		x.BigIntID, err = strconv.ParseInt(x.ID, 10, 64)
// 		if err != nil {
// 			return err
// 		}
// 	}
//
// 	if x.Status != nil {
// 		x.StatusString = x.Status.String()
// 	}
//
// 	return nil
// }
//
// // AfterFind GORM hook - convert string to OrderState after fetching from DB
// func (x *Order) AfterFind(tx *gorm.DB) (err error) {
// 	x.ID = strconv.FormatInt(x.BigIntID, 10)
//
// 	if x.StatusString != "" {
// 		x.Status, err = UnmarshalOrderState(x.StatusString)
// 		if err != nil {
// 			return err
// 		}
// 	}
//
// 	return nil
// }

// NewOrderItem creates a new order item.
func NewOrderItem(menuItemID, name string, price float64, quantity int) *OrderItem {
	return &OrderItem{
		MenuItemId: menuItemID,
		Quantity:   int64(quantity),
		Price:      price,
	}
}
