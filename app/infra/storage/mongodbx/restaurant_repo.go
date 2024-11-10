package mongodbx

import (
	"context"
	"fmt"
	"time"

	"github.com/blackhorseya/godine/app/infra/otelx"
	"github.com/blackhorseya/godine/entity/domain/restaurant/model"
	"github.com/blackhorseya/godine/entity/domain/restaurant/repo"
	"github.com/blackhorseya/godine/pkg/contextx"
	"github.com/blackhorseya/godine/pkg/persistence"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type mongodbRestaurantRepo struct {
	persistence.IRepository[*model.Restaurant]

	rw   *mongo.Client
	db   *mongo.Database
	coll *mongo.Collection
}

// NewRestaurantRepo is to create a new mongodbRestaurantRepo.
func NewRestaurantRepo(rw *mongo.Client) repo.IRestaurantRepo {
	db := rw.Database("godine")
	coll := db.Collection("restaurants")

	return &mongodbRestaurantRepo{
		IRepository: persistence.NewMongoRepository[*model.Restaurant](coll),
		rw:          rw,
		db:          db,
		coll:        coll,
	}
}

func (x *mongodbRestaurantRepo) CreateReservation(
	c context.Context,
	restaurant *model.Restaurant,
	reservation *model.Order,
) error {
	_, span := otelx.Tracer.Start(c, "mongodbx.restaurant_repo.CreateReservation")
	defer span.End()

	ctx := contextx.WithContextx(c)

	now := timestamppb.Now()
	reservation.Id = primitive.NewObjectID().Hex()
	reservation.CreatedAt = now
	reservation.UpdatedAt = now

	timeout, cancelFunc := context.WithTimeout(ctx, 5*time.Second)
	defer cancelFunc()

	// Insert the reservation into the "orders" collection
	ordersCollection := x.db.Collection("orders")
	_, err := ordersCollection.InsertOne(timeout, reservation)
	if err != nil {
		return fmt.Errorf("failed to create reservation: %w", err)
	}

	// Optionally update the restaurant document with the new reservation if it’s required
	restaurantsCollection := x.db.Collection("restaurants")
	filter := bson.M{"_id": restaurant.Id}
	update := bson.M{
		"$push": bson.M{
			"orders": reservation.Id, // Assumes orders are stored as a list of order IDs in the restaurant document
		},
	}

	_, err = restaurantsCollection.UpdateOne(timeout, filter, update)
	if err != nil {
		return fmt.Errorf("無法更新餐廳以新增訂單: %w", err)
	}

	// 庫存更新：根據訂單中的 dishes，減少對應 MenuItem 的庫存數量
	menuItemsCollection := x.db.Collection("menu_items")
	for _, dish := range reservation.Dishes {
		// 查找並更新指定的 MenuItem 庫存
		menuItemFilter := bson.M{"_id": dish.MenuItemId}
		menuItemUpdate := bson.M{
			"$inc": bson.M{
				"quantity": -int32(dish.Quantity), // 減少菜單項目的庫存
			},
		}

		_, err = menuItemsCollection.UpdateOne(timeout, menuItemFilter, menuItemUpdate)
		if err != nil {
			return fmt.Errorf("無法更新菜單項目庫存: %w", err)
		}
	}

	return nil
}
