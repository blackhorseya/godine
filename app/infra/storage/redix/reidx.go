package redix

import (
	"fmt"

	"github.com/blackhorseya/godine/app/infra/configx"
	"github.com/blackhorseya/godine/pkg/contextx"
	redisstd "github.com/redis/go-redis/v9"
	"github.com/testcontainers/testcontainers-go/modules/redis"
)

// NewClient is a function to create a new redis client
func NewClient(app *configx.Application) (*redisstd.Client, error) {
	return redisstd.NewClient(&redisstd.Options{
		Addr:     app.Storage.Redis.Addr,
		Password: "",
		DB:       0,
	}), nil
}

// Container is used to represent a redis container.
type Container struct {
	*redis.RedisContainer
}

// NewContainer returns a new redis container.
func NewContainer(ctx contextx.Contextx) (*Container, error) {
	container, err := redis.Run(
		ctx,
		"docker.io/redis:7",
		redis.WithSnapshotting(10, 1),
		redis.WithLogLevel(redis.LogLevelVerbose),
	)
	if err != nil {
		return nil, fmt.Errorf("run redis container: %w", err)
	}

	return &Container{
		RedisContainer: container,
	}, nil
}

// RW returns a read-write client.
func (c *Container) RW(ctx contextx.Contextx) (*redisstd.Client, error) {
	dsn, err := c.ConnectionString(ctx)
	if err != nil {
		return nil, err
	}

	return redisstd.NewClient(&redisstd.Options{
		Addr: dsn,
	}), nil
}
