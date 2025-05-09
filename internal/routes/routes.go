package routes

import (
	"bookhub/internal/handlers"
	"bookhub/internal/middleware"
	"bookhub/internal/session"
	"net/http"

	"github.com/gorilla/mux"
)

func SetupRoutes(r *mux.Router) {
	r.HandleFunc("/books", middleware.Chain(handlers.FetchBookHandler, middleware.Logging()))
	r.HandleFunc("/books/{id}", middleware.Chain(session.AuthMiddleware(handlers.FetchBookHandler), middleware.Logging()))
	r.HandleFunc("/books/create-book", middleware.Chain(session.AuthMiddleware(handlers.CreateBookHandler), middleware.Logging()))
	r.HandleFunc("/", middleware.Chain(handlers.HomeHandler, middleware.Logging())).Methods("GET")

	// Auth
	r.HandleFunc("/login", middleware.Chain(session.LoginHandler, middleware.Logging()))
	r.HandleFunc("/logout", middleware.Chain(session.LogoutHandler, middleware.Logging()))

	// Serve CSS, JS, and images from static dir
	fs := http.FileServer(http.Dir("static/"))
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", fs))
}
