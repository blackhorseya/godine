//go:build external

package order

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type mariadbExternalTester struct {
	suite.Suite
}

func TestMariadbExternal(t *testing.T) {
	suite.Run(t, new(mariadbExternalTester))
}
