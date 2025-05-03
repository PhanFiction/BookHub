package main

import (
	"bookhub/internal/database"
	"bookhub/internal/routes"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	defer database.DB.Close()

	r := mux.NewRouter()
	routes.SetupRoutes(r)
	http.ListenAndServe(":8080", r)
}
