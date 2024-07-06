//go:build external

package restful

import (
	"testing"

	"github.com/spf13/viper"
)

func TestRun(t *testing.T) {
	restful, err := New(viper.New())
	if err != nil {
		t.Fatalf("New() error = %v", err)
	}

	err = restful.Start()
	if err != nil {
		t.Fatalf("Start() error = %v", err)
	}

	err = restful.AwaitSignal()
	if err != nil {
		t.Fatalf("AwaitSignal() error = %v", err)
	}
}
