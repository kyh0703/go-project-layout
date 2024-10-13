package main

import (
	"context"
	"os/signal"
	"syscall"
)

func main() {
	// create context that listens for the interrupt signal from the OS
	ctx, stop := signal.NotifyContext(
		context.Background(),
		syscall.SIGINT,
		syscall.SIGTERM,
	)
	defer stop()

	// start app
	// a, err := (ctx)
	// if err != nil {
	// }
	// defer a.Close()

	// listen for the interrupt signal.
	<-ctx.Done()

	stop()
}
