package routes

import (
	"bookhub/internal/handlers"
	"bookhub/internal/middleware"
	"bookhub/internal/session"
	"net/http"

	"github.com/gorilla/mux"
)

func SetupRoutes(r *mux.Router) {
	// Middleware
	r.HandleFunc("/books/{id}", middleware.Chain(handlers.FetchSingleBookHandler, middleware.Logging())).Methods("GET")
	r.HandleFunc("/books/{id}", middleware.Chain(handlers.UpdateBookHandler, middleware.Logging())).Methods("PUT")
	r.HandleFunc("/books/{id}", middleware.Chain(handlers.DeleteBookHandler, middleware.Logging())).Methods("DELETE")
	r.HandleFunc("/books/create-book", middleware.Chain(session.AuthMiddleware(handlers.CreateBookHandler), middleware.Logging()))
	r.HandleFunc("/books/update-book/{id}", middleware.Chain(session.AuthMiddleware(handlers.UpdateBookHandler), middleware.Logging()))
	r.HandleFunc("/books", middleware.Chain(handlers.FetchBooksHandler, middleware.Logging()))
	r.HandleFunc("/books/save-book/{id}", handlers.SaveBookHandler)
	r.HandleFunc("/", middleware.Chain(handlers.HomeHandler, middleware.Logging())).Methods("GET")

	// Auth
	r.HandleFunc("/login", middleware.Chain(session.LoginHandler, middleware.Logging()))
	r.HandleFunc("/signup", middleware.Chain(session.SignupHandler, middleware.Logging()))
	r.HandleFunc("/logout", middleware.Chain(session.LogoutHandler, middleware.Logging()))
	r.HandleFunc("/update-user", middleware.Chain(session.AuthMiddleware(session.UpdateUserHandler), middleware.Logging()))
	r.HandleFunc("/google-auth", middleware.Chain(handlers.GoogleAuthHandler, middleware.Logging())).Methods("POST")

	// Serve CSS, JS, and images from static dir
	fs := http.FileServer(http.Dir("static/"))
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", fs))
}
