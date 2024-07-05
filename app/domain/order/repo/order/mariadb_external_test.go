//go:build external

package order

import (
	"testing"

	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/suite"
)

type mariadbExternalTester struct {
	suite.Suite

	rw *sqlx.DB
}

func (s *mariadbExternalTester) SetupTest() {
}

func (s *mariadbExternalTester) TearDownTest() {
	if s.rw != nil {
		_ = s.rw.Close()
	}
}

func TestMariadbExternal(t *testing.T) {
	suite.Run(t, new(mariadbExternalTester))
}
