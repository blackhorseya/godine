//go:build external

package platform

import (
	"os"
	"os/signal"
	"syscall"
	"testing"

	"github.com/blackhorseya/godine/pkg/contextx"
	"github.com/spf13/viper"
)

func TestRun(t *testing.T) {
	service, clean, err := New(viper.GetViper())
	if err != nil {
		t.Fatalf("New() error = %v", err)
	}
	defer clean()

	ctx := contextx.Background()

	err = service.Start(ctx)
	if err != nil {
		t.Fatalf("Start() error = %v", err)
	}

	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, syscall.SIGTERM, syscall.SIGINT)

	<-signalChan

	err = service.Shutdown(ctx)
	if err != nil {
		t.Fatalf("AwaitSignal() error = %v", err)
	}
}
