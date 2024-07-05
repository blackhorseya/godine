package mariadbx

import (
	"github.com/blackhorseya/godine/app/infra/configx"
	_ "github.com/go-sql-driver/mysql" // import MySQL driver
	"github.com/jmoiron/sqlx"
)

// NewClient init mysql client.
func NewClient(app *configx.Application) (*sqlx.DB, error) {
	// todo: 2024/7/5|sean|implement me
	panic("implement me")
}
