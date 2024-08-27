//go:build integration

package biz

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type suiteRestaurantServiceIntegration struct {
	suite.Suite
}

func TestIntegrationAll(t *testing.T) {
	suite.Run(t, new(suiteRestaurantServiceIntegration))
}
