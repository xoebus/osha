package shutdown

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
)

// The signals which mean that the program should shut down.
var signals = []os.Signal{os.Interrupt, syscall.SIGTERM}

// WithShutdown returns an enriched version of the original Context which will
// be cancelled when the user wants to shut down the program by sending SIGINT
// or SIGTERM. If the user sends the signal again before the program has
// finished shutting down then it will exit immediately with exit status 1.
func WithShutdown(ctx context.Context) context.Context {
	sctx, cancel := context.WithCancel(ctx)
	c := make(chan os.Signal, 2)
	signal.Notify(c, signals...)
	go func() {
		<-c
		cancel()
		<-c
		log.Println("user interrupted again; shutting down immediately...")
		os.Exit(1)
	}()
	return sctx
}
