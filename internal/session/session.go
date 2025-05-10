package session

import (
	"bookhub/internal/auth"
	"bookhub/internal/database"
	"bookhub/internal/types"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/sessions"
)

// Global var to be used
// key must be 16, 24 or 32 bytes long (AES-128, AES-192 or AES-256)
var Store = sessions.NewCookieStore([]byte("super-secret-key"))

// Handles login for user
func LoginHandler(w http.ResponseWriter, r *http.Request) {
	var body types.User
	json.NewDecoder(r.Body).Decode(&body)

	// Check if user exists in the database
	userData, err := database.GetUser(database.DB, body.Username)
	fmt.Println(userData.Username, userData.Password)

	if err != nil {
		http.Error(w, "Could not retrieve user", http.StatusInternalServerError)
		return
	}

	isValid := auth.CheckPasswordHash(body.Password, userData.Password)

	if !isValid {
		http.Error(w, "Invalid credentials", http.StatusUnauthorized)
		return
	}

	session, _ := Store.Get(r, "session")
	// Authentication goes here
	// Set user as authenticated
	session.Values["authenticated"] = true
	session.Values["user_id"] = userData.ID
	session.Values["username"] = userData.Username
	session.Values["email"] = userData.Email
	session.Save(r, w)

	// Send JSON Data
	json.NewEncoder(w).Encode(map[string]interface{}{
		"message":       "Login successful",
		"authenticated": true,
		"user_id":       userData.ID,
		"username":      userData.Username,
	})
}

// Handles the signup process
func SignupHandler(w http.ResponseWriter, r *http.Request) {
	var body types.User
	json.NewDecoder(r.Body).Decode(&body)

	// Save user to database
	// Hash password
	hashedPassword, err := auth.HashPassword(body.Password)
	if err != nil {
		http.Error(w, "Could not hash password", http.StatusInternalServerError)
		return
	}

	// Save user to database
	database.CreateUser(database.DB, body.Email, body.Name, body.Username, hashedPassword)

	// Redirect to login
	http.Redirect(w, r, "/login", http.StatusSeeOther)
}

func UpdateUserHandler(w http.ResponseWriter, r *http.Request) {
	session, _ := Store.Get(r, "session")
	userID, _ := session.Values["user_id"].(int)

	var body types.User
	json.NewDecoder(r.Body).Decode(&body)

	hashedPassword, password_err := auth.HashPassword(body.Password)
	if password_err != nil {
		fmt.Println("Could not hash password")
		return
	}

	userIdString := strconv.Itoa(userID)

	err := database.UpdateUser(database.DB, userIdString, body.Name, body.Username, body.Email, hashedPassword)

	if err != nil {
		http.Error(w, "Could not update account.", http.StatusInternalServerError)
		return
	}

	session.Values["username"] = body.Username
	session.Values["email"] = body.Email

	json.NewEncoder(w).Encode(map[string]interface{}{
		"message":       "Login successful",
		"authenticated": true,
		"user_id":       userID,
		"username":      body.Username,
	})
}

// Handles the logout process
func LogoutHandler(w http.ResponseWriter, r *http.Request) {
	session, _ := Store.Get(r, "session")

	// Revoke users authentication
	session.Values["authenticated"] = false
	session.Save(r, w)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"message":       "Logged out successful",
		"authenticated": false,
	})
}

// Middleware to check if user is authenticated
func AuthMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		session, _ := Store.Get(r, "session")
		auth, ok := session.Values["authenticated"].(bool)
		if !ok || !auth {
			http.Redirect(w, r, "/", http.StatusSeeOther)
			return
		}
		next(w, r)
	}
}

func CreateFakeUser() {
	hashedPassword, err := auth.HashPassword("tester")
	if err != nil {
		fmt.Println("Could not hash password")
		return
	}

	// Save user to database
	database.CreateUser(database.DB, "tester@gmail.com", "Tester", "tester", hashedPassword)
	// db.Exec("INSERT INTO users (username, password) VALUES ($1, $2)", username, hashedPassword)
	// For now, just print to console
	fmt.Printf("User %s with password %s saved to database\n", "tester", hashedPassword)
}
