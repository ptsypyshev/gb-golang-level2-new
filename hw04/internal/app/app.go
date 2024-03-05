package app

import (
	"context"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/ptsypyshev/gb-golang-level2-new/hw04/internal/repo"
	"github.com/ptsypyshev/gb-golang-level2-new/hw04/internal/usecase"
)

// App is a main application struct.
type App struct {
	router  *chi.Mux
	users   repo.UserRepo
	friendship usecase.Friendship
}

// New is a constructor for App.
func New(u repo.UserRepo, f usecase.Friendship) *App {
	r := chi.NewRouter()
	return &App{
		router: r,
		users: u,
		friendship: f,
	}
}

// Run is a main method to start serve connections.
func (a *App) Run(ctx context.Context) {
	server := &http.Server{Addr: ":8000", Handler: a.router}
	go server.ListenAndServe()
	<- ctx.Done()
	server.Shutdown(context.Background())
}
