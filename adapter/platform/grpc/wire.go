//go:build wireinject

//go:generate wire

package grpc

import (
	"github.com/blackhorseya/godine/pkg/adapterx"
	"github.com/google/wire"
	"github.com/spf13/viper"
)

func New(v *viper.Viper) (adapterx.Restful, error) {
	panic(wire.Build(NewServer))
}
