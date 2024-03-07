package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/ptsypyshev/gb-golang-level2-new/hw04/internal/app"
	"github.com/ptsypyshev/gb-golang-level2-new/hw04/internal/repo/friends"
	"github.com/ptsypyshev/gb-golang-level2-new/hw04/internal/repo/users"
	"github.com/ptsypyshev/gb-golang-level2-new/hw04/internal/storage/inmem"
	"github.com/ptsypyshev/gb-golang-level2-new/hw04/internal/usecase/friendship"
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

	stor := inmem.New()
	uRepo := users.New(stor)
	fRepo := friends.New(stor)
	fUsecase := friendship.New(uRepo, fRepo)

	a := app.New(uRepo, fUsecase)
	a.SetupRoutes()
	a.Run(ctx)
}
