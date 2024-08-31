package postgresqlx

import (
	"fmt"
	"time"

	"github.com/blackhorseya/godine/app/infra/configx"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

const (
	defaultConns       = 100
	defaultMaxLifetime = 15 * time.Minute

	defaultTimeout  = 5 * time.Second
	defaultLimit    = 10
	defaultMaxLimit = 100
)

// NewClient init mysql client.
func NewClient(app *configx.Application) (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(app.Storage.Postgresql.DSN), &gorm.Config{
		SkipDefaultTransaction: true,
		PrepareStmt:            true,
	})
	if err != nil {
		return nil, fmt.Errorf("open postgres client error: %w", err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		return nil, fmt.Errorf("get postgres db error: %w", err)
	}

	sqlDB.SetConnMaxLifetime(defaultMaxLifetime)
	sqlDB.SetMaxOpenConns(defaultConns)
	sqlDB.SetMaxIdleConns(defaultConns)

	// register custom data type
	schema.RegisterSerializer("timestamppb", TimestampSerializer{})

	return db, nil
}
