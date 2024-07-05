package mariadbx

import (
	"fmt"
	"time"

	"github.com/blackhorseya/godine/app/infra/configx"
	_ "github.com/go-sql-driver/mysql" // import MySQL driver
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/mysql"
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

// AutoMigrate auto migrate the database.
func AutoMigrate(db *sqlx.DB, source string, dbName string) error {
	driver, err := mysql.WithInstance(db.DB, &mysql.Config{DatabaseName: dbName})
	if err != nil {
		return fmt.Errorf("create mysql driver error: %w", err)
	}

	m, err := migrate.NewWithDatabaseInstance(source, "mysql", driver)
	if err != nil {
		return fmt.Errorf("create migration instance error: %w", err)
	}

	return m.Up()
}
