package handlers

import (
	"bookhub/internal/session"
	"bookhub/internal/types"
	"encoding/json"
	"net/http"
)

// Return single book from the database
func FetchBookHandler(w http.ResponseWriter, r *http.Request) {
	session, _ := session.Store.Get(r, "session")
	auth, ok := session.Values["authenticated"].(bool)

	data := types.Data{
		BookData: types.BookDetails{
			Title:       "The Great Gatsby",
			Author:      "F. Scott Fitzgerald",
			Pages:       180,
			Publisher:   "Scribner",
			ISBN:        "9780743273565",
			Description: "A novel about the American dream.",
			PublishedAt: "1925-04-10",
		},
		Authenticated: ok && auth,
	}
	json.NewEncoder(w).Encode(data)
}

func CreateBookHandler(w http.ResponseWriter, r *http.Request) {
	session, _ := session.Store.Get(r, "session")
	auth, ok := session.Values["authenticated"].(bool)

	data := types.Data{
		Authenticated: ok && auth,
	}
	json.NewEncoder(w).Encode(data)
}
