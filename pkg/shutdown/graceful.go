package shutdown

import (
	"os"
	"os/signal"
	"syscall"
)

// Graceful waits for a signal to shutdown the application.
func Graceful() {
	quit := make(chan os.Signal, 1)
	defer close(quit)

	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
}
