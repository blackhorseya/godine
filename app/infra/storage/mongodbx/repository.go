package mongodbx

import (
	"go.mongodb.org/mongo-driver/mongo"
)

// MongoRepository is a struct for mongo repository.
type MongoRepository struct {
	coll *mongo.Collection
}
