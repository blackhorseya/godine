package redix

import (
	"github.com/blackhorseya/godine/app/infra/configx"
	"github.com/redis/go-redis/v9"
)

// NewRedis is a function to create a new redis client
func NewRedis(app *configx.Application) (*redis.Client, error) {
	return redis.NewClient(&redis.Options{
		Addr:     app.Storage.Redis.Addr,
		Password: "",
		DB:       0,
	}), nil
}
