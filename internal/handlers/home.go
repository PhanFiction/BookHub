package handlers

import (
	"bookhub/internal/session"
	"bookhub/internal/types"
	"encoding/json"
	"net/http"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	session, _ := session.Store.Get(r, "session")
	auth, ok := session.Values["authenticated"].(bool)

	data := types.Data{
		Authenticated: ok && auth,
	}
	json.NewEncoder(w).Encode(data)
}
