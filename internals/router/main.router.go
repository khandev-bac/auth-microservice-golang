package router

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/khandev-bac/lemon/config"
	"github.com/khandev-bac/lemon/internals/firebase"
	"github.com/khandev-bac/lemon/internals/handler"
	"github.com/khandev-bac/lemon/internals/middleware"
	"github.com/khandev-bac/lemon/internals/repo"
	"github.com/khandev-bac/lemon/internals/service"
)

func MainRoute() http.Handler {
	firebasePath := firebase.NewFirebaseService(config.AppConfig.FirebasePath)
	db := config.DB
	repo := repo.NewRepo(db)
	service := service.NewService(repo)
	handler := handler.NewHandler(service, firebasePath)

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
