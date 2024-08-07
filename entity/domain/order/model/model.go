package model

import (
	"encoding/json"
	"strconv"
	"time"

	"github.com/blackhorseya/godine/pkg/contextx"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"gorm.io/gorm"
)

// Order represents an order entity.
type Order struct {
	// ID is the unique identifier of the order.
	ID string `json:"id,omitempty" bson:"_id,omitempty" gorm:"-"`

	// BigIntID is the actual field stored as BIGINT in the DB.
	BigIntID int64 `json:"-" gorm:"column:id;primaryKey;not null"`

	// UserID is the identifier of the user who placed the order.
	UserID string `json:"user_id,omitempty" bson:"user_id" gorm:"column:user_id;not null"`

	// RestaurantID is the identifier of the restaurant where the order was placed.
	RestaurantID string `json:"restaurant_id,omitempty" bson:"restaurant_id" gorm:"column:restaurant_id;not null"`

	// Items are the list of items in the order.
	Items []OrderItem `json:"items,omitempty" bson:"items" gorm:"foreignKey:BigIntOrderID;references:BigIntID"`

	// Status is the current status of the order (e.g., pending, confirmed, delivered).
	Status       OrderState `json:"status,omitempty" bson:"status" gorm:"-"`
	StatusString string     `json:"-" bson:"-" gorm:"column:status;not null"`

	// TotalAmount is the total amount of the order.
	TotalAmount float64 `json:"total_amount,omitempty" bson:"total_amount" gorm:"column:total_amount"`

	// CreatedAt is the timestamp when the order was created.
	CreatedAt time.Time `json:"created_at,omitempty" bson:"created_at" gorm:"column:created_at;autoCreateTime"`

	// UpdatedAt is the timestamp when the order was last updated.
	UpdatedAt time.Time `json:"updated_at,omitempty" bson:"updated_at" gorm:"column:updated_at;autoUpdateTime"`

	// DeliveryID is the identifier of the delivery associated with the order.
	DeliveryID string `json:"delivery_id,omitempty" bson:"delivery_id" gorm:"column:delivery_id"`
}

func (x *Order) MarshalJSON() ([]byte, error) {
	type Alias Order
	alias := &struct {
		*Alias `json:",inline"`
		Status string `json:"status,omitempty"`
	}{
		Alias:  (*Alias)(x),
		Status: "",
	}

	if x.Status != nil {
		alias.Status = x.Status.String()
	}

	return json.Marshal(alias)
}

func (x *Order) UnmarshalBSON(bytes []byte) error {
	type Alias Order
	alias := &struct {
		ID     primitive.ObjectID `bson:"_id"`
		Status string             `bson:"status"`
		*Alias `bson:",inline"`
	}{
		Alias: (*Alias)(x),
	}

	if err := bson.Unmarshal(bytes, alias); err != nil {
		return err
	}

	x.ID = alias.ID.Hex()

	state, err := UnmarshalOrderState(alias.Status)
	if err != nil {
		return err
	}
	x.Status = state

	return nil
}

func (x *Order) MarshalBSON() ([]byte, error) {
	type Alias Order
	alias := &struct {
		*Alias `bson:",inline"`
		Status string             `bson:"status"`
		ID     primitive.ObjectID `bson:"_id"`
	}{
		Alias:  (*Alias)(x),
		Status: x.Status.String(),
	}

	id, err := primitive.ObjectIDFromHex(x.ID)
	if err != nil {
		return nil, err
	}
	alias.ID = id

	return bson.Marshal(alias)
}

// NewOrder creates a new order.
func NewOrder(userID, restaurantID string, items []OrderItem) *Order {
	totalAmount := 0.0
	for _, item := range items {
		totalAmount += item.Price * float64(item.Quantity)
	}

	return &Order{
		UserID:       userID,
		RestaurantID: restaurantID,
		Items:        items,
		Status:       &PendingState{},
		TotalAmount:  totalAmount,
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
		DeliveryID:   "",
	}
}

// Next transitions the order to the next state.
func (x *Order) Next(ctx contextx.Contextx) (event *OrderEvent, err error) {
	return x.Status.Next(ctx, x)
}

// AddItem adds an item to the order.
func (x *Order) AddItem(item OrderItem) {
	x.Items = append(x.Items, item)
	x.UpdatedAt = time.Now()
}

// BeforeSave GORM hook - convert OrderState to string before saving
func (x *Order) BeforeSave(tx *gorm.DB) (err error) {
	if x.ID != "" {
		x.BigIntID, err = strconv.ParseInt(x.ID, 10, 64)
		if err != nil {
			return err
		}
	}

	if x.Status != nil {
		x.StatusString = x.Status.String()
	}

	return nil
}

// AfterFind GORM hook - convert string to OrderState after fetching from DB
func (x *Order) AfterFind(tx *gorm.DB) (err error) {
	x.ID = strconv.FormatInt(x.BigIntID, 10)

	if x.StatusString != "" {
		x.Status, err = UnmarshalOrderState(x.StatusString)
		if err != nil {
			return err
		}
	}

	return nil
}

// OrderItem represents an item in the order.
type OrderItem struct {
	// OrderID is the identifier of the order to which the item belongs.
	OrderID string `json:"order_id,omitempty" bson:"order_id" gorm:"-"`

	// BigIntOrderID is the actual field stored as BIGINT in the DB.
	BigIntOrderID int64 `json:"-" gorm:"column:order_id;primaryKey;not null"`

	// MenuItemID is the identifier of the menu item.
	MenuItemID string `json:"menu_item_id,omitempty" bson:"menu_item_id" gorm:"column:item_id;primaryKey"`

	// Quantity is the quantity of the menu item ordered.
	Quantity int `json:"quantity,omitempty" bson:"quantity" example:"2" gorm:"column:quantity"`

	// Price is the price of a single unit of the menu item.
	Price float64 `json:"price,omitempty" bson:"price" example:"10.5" gorm:"column:price"`
}

// NewOrderItem creates a new order item.
func NewOrderItem(menuItemID, name string, price float64, quantity int) *OrderItem {
	return &OrderItem{
		MenuItemID: menuItemID,
		Quantity:   quantity,
		Price:      price,
	}
}
