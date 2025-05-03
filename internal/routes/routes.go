package routes

import (
	"bookhub/internal/handlers"
	"bookhub/internal/middleware"
	"bookhub/internal/session"
	"net/http"

	"github.com/gorilla/mux"
)

func SetupRoutes(r *mux.Router) {
	r.HandleFunc("/books/{title}/id/{id}", middleware.Chain(handlers.FetchBook, middleware.Logging())).Methods("GET")
	r.HandleFunc("/create-book", middleware.Chain(handlers.CreateBook, middleware.Logging())).Methods("POST")
	r.HandleFunc("/", middleware.Chain(handlers.Home, middleware.Logging())).Methods("GET")

	// Auth
	r.HandleFunc("/login", middleware.Chain(session.LoginHandler, middleware.Logging()))
	r.HandleFunc("/logout", middleware.Chain(session.LogoutHandler, middleware.Logging()))

	// Serve CSS, JS, and images from static dir
	fs := http.FileServer(http.Dir("static/"))
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", fs))
}
