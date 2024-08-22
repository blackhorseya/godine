//go:build external

package platform

import (
	"os"
	"os/signal"
	"syscall"
	"testing"

	"github.com/spf13/viper"
)

func TestRun(t *testing.T) {
	service, err := New(viper.New())
	if err != nil {
		t.Fatalf("New() error = %v", err)
	}

	err = service.Start()
	if err != nil {
		t.Fatalf("Start() error = %v", err)
	}

	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, syscall.SIGTERM, syscall.SIGINT)

	<-signalChan

	err = service.AwaitSignal()
	if err != nil {
		t.Fatalf("AwaitSignal() error = %v", err)
	}
}
