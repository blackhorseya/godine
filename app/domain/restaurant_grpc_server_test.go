package domain

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type suiteRestaurantServiceTester struct {
	suite.Suite
}

func TestAllRestaurantServiceTester(t *testing.T) {
	suite.Run(t, new(suiteRestaurantServiceTester))
}
