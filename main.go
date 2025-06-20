package main

import (
	"bookhub/internal/database"
	"bookhub/internal/routes"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func main() {
	defer database.DB.Close()

	r := mux.NewRouter() // Use Gorilla Mux router
	routes.SetupRoutes(r)

	// Enable CORS middleware
	// CORS must wrap `r` — this is what I mean by "correct handler"
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:3000"}, // Your frontend
		AllowCredentials: true,
		AllowedHeaders:   []string{"Authorization", "Content-Type"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE"},
	})

	handler := c.Handler(r)

	log.Println("Server listening on http://localhost:8080")
	http.ListenAndServe(":8080", handler)
}
