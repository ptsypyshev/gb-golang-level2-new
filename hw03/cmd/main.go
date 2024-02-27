package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/ptsypyshev/gb-golang-level2-new/hw03/internal/conveyor"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())

	go func() {
		exit := make(chan os.Signal, 1)
		signal.Notify(exit, os.Interrupt, syscall.SIGTERM)
		<-exit
		fmt.Println("Graceful shutdown by signal!")
		cancel()
	}()

	c := conveyor.New()
	c.Run(ctx)
}
