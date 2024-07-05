//go:build external

package order

import (
	"testing"

	"github.com/blackhorseya/godine/app/infra/configx"
	"github.com/blackhorseya/godine/app/infra/storage/mariadbx"
	"github.com/blackhorseya/godine/entity/domain/order/repo"
	"github.com/blackhorseya/godine/pkg/logging"
	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/suite"
)

type mariadbExternalTester struct {
	suite.Suite

	rw   *sqlx.DB
	repo repo.IOrderRepo
}

func (s *mariadbExternalTester) SetupTest() {
	err := configx.LoadConfig("")
	s.Require().NoError(err)

	app, err := configx.LoadApplication(&configx.C.OrderRestful)
	s.Require().NoError(err)

	err = logging.Init(app.Log)
	s.Require().NoError(err)

	rw, err := mariadbx.NewClient(app)
	s.Require().NoError(err)
	s.rw = rw

	s.repo = NewMariadb(s.rw)
}

func (s *mariadbExternalTester) TearDownTest() {
	if s.rw != nil {
		_ = s.rw.Close()
	}
}

func TestMariadbExternal(t *testing.T) {
	suite.Run(t, new(mariadbExternalTester))
}
