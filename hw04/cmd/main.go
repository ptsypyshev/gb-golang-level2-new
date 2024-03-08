package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/ptsypyshev/gb-golang-level2-new/hw04/internal/app"
	"github.com/ptsypyshev/gb-golang-level2-new/hw04/internal/repo/friends"
	"github.com/ptsypyshev/gb-golang-level2-new/hw04/internal/repo/users"
	"github.com/ptsypyshev/gb-golang-level2-new/hw04/internal/storage/pgdb"
	"github.com/ptsypyshev/gb-golang-level2-new/hw04/internal/usecase/friendship"
)

const connStr = "postgres://postgres:postgres@pg-friends:5432/friends?sslmode=disable"

func main() {
	ctx, cancel := context.WithCancel(context.Background())

	go func() {
		exit := make(chan os.Signal, 1)
		signal.Notify(exit, os.Interrupt, syscall.SIGTERM)
		<-exit
		fmt.Println("Graceful shutdown by signal!")
		cancel()
	}()

	stor, err := pgdb.New(connStr)
	if err != nil {
		log.Panicf("cannot init DB: %s", err)
	}
	defer stor.Close()

	uRepo := users.New(stor)
	fRepo := friends.New(stor)
	fUsecase := friendship.New(uRepo, fRepo)

	a := app.New(uRepo, fUsecase)
	a.SetupRoutes()
	a.Run(ctx)
}
