package mariadbx

import (
	"fmt"
	"time"

	"github.com/blackhorseya/godine/app/infra/configx"
	_ "github.com/go-sql-driver/mysql" // import MySQL driver
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

const (
	defaultConns       = 100
	defaultMaxLifetime = 15 * time.Minute
)

// NewClient init mysql client.
func NewClient(app *configx.Application) (*gorm.DB, error) {
	db, err := gorm.Open(mysql.Open(app.Storage.Mysql.DSN), &gorm.Config{
		SkipDefaultTransaction: true,
		PrepareStmt:            true,
	})
	if err != nil {
		return nil, fmt.Errorf("open mysql client error: %w", err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		return nil, fmt.Errorf("get mysql db error: %w", err)
	}

	sqlDB.SetConnMaxLifetime(defaultMaxLifetime)
	sqlDB.SetMaxOpenConns(defaultConns)
	sqlDB.SetMaxIdleConns(defaultConns)

	return db, nil
}
