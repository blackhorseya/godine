package mariadbx

import (
	"fmt"
	"time"

	"github.com/blackhorseya/godine/app/infra/configx"
	_ "github.com/go-sql-driver/mysql" // import MySQL driver
	"github.com/jmoiron/sqlx"
)

const (
	defaultConns       = 100
	defaultMaxLifetime = 15 * time.Minute
)

// NewClient init mysql client.
func NewClient(app *configx.Application) (*sqlx.DB, error) {
	db, err := sqlx.Open("mysql", app.Storage.Mysql.DSN)
	if err != nil {
		return nil, fmt.Errorf("open mysql client error: %w", err)
	}

	db.SetConnMaxLifetime(defaultMaxLifetime)
	db.SetMaxOpenConns(defaultConns)
	db.SetMaxIdleConns(defaultConns)

	return db, nil
}
