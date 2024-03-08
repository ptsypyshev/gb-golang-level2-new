package app

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

// SetupRoutes defines all routers and their handlers.
func (a *App) SetupRoutes() {
	a.router.Use(middleware.Recoverer)
	a.router.Route("/users", func(r chi.Router) {
		r.Get("/{user_id}", a.ReadUser)
		r.Post("/", a.CreateUser)
		r.Put("/", a.UpdateUser)
		r.Delete("/", a.DeleteUser)
	})
	a.router.Route("/friends", func(r chi.Router) {
		r.Get("/{user_id}", a.GetUserFriends)
		r.Post("/", a.CreateFriendship)
	})
}
