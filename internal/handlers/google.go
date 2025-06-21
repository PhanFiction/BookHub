package handlers

import (
	"bookhub/internal/database"
	"bookhub/internal/session"
	"bookhub/internal/types"
	"context"
	"encoding/json"
	"net/http"
	"os"

	"google.golang.org/api/idtoken"
)

func GoogleAuthHandler(w http.ResponseWriter, r *http.Request) {
	var body struct {
		Token string `json:"token"`
	}

	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	// Validate the token using Google's ID token package
	payload, err := idtoken.Validate(context.Background(), body.Token, os.Getenv("GOOGLE_CLIENT_ID"))
	if err != nil {
		http.Error(w, "Invalid Google token", http.StatusUnauthorized)
		return
	}

	googleID := payload.Subject
	email, _ := payload.Claims["email"].(string)
	name, _ := payload.Claims["name"].(string)
	givenName, _ := payload.Claims["given_name"].(string)
	familyName, _ := payload.Claims["family_name"].(string)
	avatar, _ := payload.Claims["picture"].(string)

	username := givenName + familyName

	userDetail := types.User{
		GoogleID: googleID,
		Email:    email,
		Username: username,
		Name:     name,
		Avatar:   avatar,
	}

	// Check if user already exists in the database
	user, err := database.GoogleAuth(database.DB, userDetail)

	if err != nil {
		http.Error(w, "Database error", http.StatusInternalServerError)
		return
	}

	// Save user session or generate a JWT
	session, _ := session.Store.Get(r, "session")
	session.Values["authenticated"] = true
	session.Values["user_id"] = user.ID
	session.Values["username"] = user.Username
	session.Save(r, w)

	json.NewEncoder(w).Encode(struct {
		Message string     `json:"message"`
		User    types.User `json:"user"`
	}{
		Message: "Logged in successfully",
		User:    user,
	})
}
