package router

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/khandev-bac/lemon/internals/handler"
	"github.com/khandev-bac/lemon/internals/middleware"
)

func MainRoute(handler *handler.UserHandler) http.Handler {
	r := chi.NewRouter()
	//public routes
	r.Post("/sign-up", handler.Signup)
	r.Post("/login", handler.Login)
	r.Post("/firebase-login", handler.FirebaseLogin)
	r.Post("/refresh-token", handler.RefreshToken)
	// private routes
	r.Group(func(r chi.Router) {
		r.Use(middleware.Auth)
		r.Post("/logout", handler.Logout)
		r.Delete("/delete-account", handler.DeleteAccount)
	})
	return r
}
