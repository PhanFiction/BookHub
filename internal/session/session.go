package session

import (
	"bookhub/internal/auth"
	"bookhub/internal/database"
	"net/http"

	"github.com/gorilla/sessions"
)

// Global var to be used
// key must be 16, 24 or 32 bytes long (AES-128, AES-192 or AES-256)
var Store = sessions.NewCookieStore([]byte("super-secret-key"))

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	username := r.FormValue("username")
	password := r.FormValue("password")

	if username == "admin" && password == "password" {
		session, _ := Store.Get(r, "session")
		// Authentication goes here
		// Set user as authenticated
		session.Values["authenticated"] = true
		session.Save(r, w)
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	http.Error(w, "Invalid credentials", http.StatusUnauthorized)
}

// Handles the logout process
func LogoutHandler(w http.ResponseWriter, r *http.Request) {
	session, _ := Store.Get(r, "session")

	// Revoke users authentication
	session.Values["authenticated"] = false
	session.Save(r, w)
	http.Redirect(w, r, "/login", http.StatusSeeOther)
}

// Middleware to check if user is authenticated
func AuthMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		session, _ := Store.Get(r, "session")
		auth, ok := session.Values["authenticated"].(bool)
		if !ok || !auth {
			http.Redirect(w, r, "/login", http.StatusSeeOther)
			return
		}
		next(w, r)
	}
}

func SignupHandler(w http.ResponseWriter, r *http.Request) {
	username := r.FormValue("username")
	name := r.FormValue("name")
	email := r.FormValue("email")
	password := r.FormValue("password")

	// Save user to database
	// Hash password
	hashedPassword, err := auth.HashPassword(password)
	if err != nil {
		http.Error(w, "Could not hash password", http.StatusInternalServerError)
		return
	}

	// Save user to database
	database.SaveUser(database.DB, email, name, username, hashedPassword)
	// db.Exec("INSERT INTO users (username, password) VALUES ($1, $2)", username, hashedPassword)
	// For now, just print to console
	// fmt.Printf("User %s with password %s saved to database\n", username, hashedPassword)

	// Redirect to login
	http.Redirect(w, r, "/login", http.StatusSeeOther)
}
