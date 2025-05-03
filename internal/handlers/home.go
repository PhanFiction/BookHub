package handlers

import (
	"bookhub/internal/session"
	"encoding/json"
	"net/http"
)

func Home(w http.ResponseWriter, r *http.Request) {
	session, _ := session.Store.Get(r, "session")
	auth, ok := session.Values["authenticated"].(bool)

	data := types.PageData{
		TabTitle:      "About Page",
		PageTitle:     "About",
		Authenticated: ok && auth,
	}
	json.NewEncoder(w).Encode(data)
}
