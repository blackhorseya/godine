package postgresqlx

import (
	"fmt"
	"time"

	"github.com/blackhorseya/godine/app/infra/configx"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const (
	defaultConns       = 100
	defaultMaxLifetime = 15 * time.Minute
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

	return db, nil
}
