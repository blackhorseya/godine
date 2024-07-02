package mongodbx

import (
	"time"

	"github.com/blackhorseya/godine/app/infra/configx"
	"github.com/blackhorseya/godine/pkg/contextx"
	"github.com/pkg/errors"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/modules/mongodb"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// NewClientWithDSN returns a new mongo client with dsn.
func NewClientWithDSN(dsn string) (*mongo.Client, error) {
	opts := options.Client().ApplyURI(dsn).
		SetMaxPoolSize(500).
		SetMinPoolSize(10).
		SetMaxConnIdleTime(10 * time.Minute).
		SetConnectTimeout(10 * time.Second).
		SetRetryWrites(true).
		SetServerSelectionTimeout(5 * time.Second)

	client, err := mongo.Connect(contextx.Background(), opts)
	if err != nil {
		return nil, err
	}

	return client, nil
}

// NewClient returns a new mongo client.
func NewClient(app *configx.Application) (*mongo.Client, error) {
	return NewClientWithDSN(app.Storage.Mongodb.DSN)
}

// Container is used to represent a mongodb container.
type Container struct {
	*mongodb.MongoDBContainer
}

// NewContainer returns a new mongodb container.
func NewContainer(ctx contextx.Contextx) (*Container, error) {
	container, err := mongodb.RunContainer(ctx, testcontainers.WithImage("mongo:6"))
	if err != nil {
		return nil, errors.Wrap(err, "run mongodb container")
	}

	return &Container{
		MongoDBContainer: container,
	}, nil
}
