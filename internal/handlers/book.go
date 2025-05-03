package handlers

import (
	"bookhub/internal/session"
	"encoding/json"
	"net/http"
)

// Return single book from the database
func FetchBook(w http.ResponseWriter, r *http.Request) {
	session, _ := session.Store.Get(r, "session")
	auth, ok := session.Values["authenticated"].(bool)

	data := types.PageData{
		TabTitle:      "About Page",
		PageTitle:     "About",
		Authenticated: ok && auth,
	}
	json.NewEncoder(w).Encode(data)
}

func CreateBook(w http.ResponseWriter, r *http.Request) {
	session, _ := session.Store.Get(r, "session")
	auth, ok := session.Values["authenticated"].(bool)

	data := types.PageData{
		TabTitle:      "About Page",
		PageTitle:     "About",
		Authenticated: ok && auth,
	}
	json.NewEncoder(w).Encode(data)
}

// Return Book based on Genre from the database

// Return Book by author

// Query Book in database
