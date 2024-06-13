package order

import (
	"go.mongodb.org/mongo-driver/mongo"
)

type mongodb struct {
	rw *mongo.Client
}
