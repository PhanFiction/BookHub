package database

import (
	"bookhub/internal/types"
	"database/sql"
	"fmt"
	"log"
	"strings"

	_ "github.com/lib/pq"
)

func CreateUserTable(db *sql.DB) {
	query := `
	CREATE TABLE IF NOT EXISTS users (
		id INTEGER PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
		google_id TEXT UNIQUE,
		password Text,
		email TEXT UNIQUE NOT NULL,
		name TEXT,
		username TEXT UNIQUE NOT NULL,
		avatar_url TEXT,
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
	);
	`
	_, err := db.Exec(query)

	if err != nil {
		log.Fatal("Error creating users table:", err)
	}

	fmt.Println("Users table created or already exists.")
}

func GetUser(db *sql.DB, email string) (types.User, error) {
	query := `SELECT id, username, name, email, password FROM users WHERE email = $1;`
	var user types.User

	// return single row
	err := db.QueryRow(query, email).Scan(&user.ID, &user.Username, &user.Name, &user.Email, &user.Password)

	fmt.Println(user, err)

	if err != nil {
		return user, err
	}

	return user, nil
}

func CreateUser(db *sql.DB, email, name, username, password string) {
	query := `
	INSERT INTO users (email, name, username, password)
	VALUES ($1, $2, $3, $4)
	ON CONFLICT DO NOTHING;
	`

	_, err := db.Exec(query, email, name, username, password)

	if err != nil {
		log.Fatal("Error saving user to database:", err)
	}

	fmt.Println("User saved to database.")
}

func UpdateUser(db *sql.DB, id string, name, username, email, password string) error {
	query := `UPDATE users SET name = $1, username = $2, email = $3, password = $4 WHERE id = $5;`
	_, err := db.Exec(query, name, username, email, password, id)

	if err != nil {
		log.Fatal("Error updating user:", err)
	}
	fmt.Println("User updated successfully.")

	return err
}

func GoogleAuth(db *sql.DB, UserDetail types.User) (types.User, error) {
	var user types.User
	username := user.FamilyName + user.FamilyName

	query := `SELECT id, google_id, email, name, username, avatar_url FROM users WHERE google_id = $1 OR email = $2;`

	err := db.QueryRow(query, UserDetail.GoogleID, UserDetail.Email).Scan(&user.ID, &user.GoogleID, &user.Email, &user.Name, &username, &user.Avatar)

	// Check if user exists
	// If user doesn't exist, create a new user
	// If user exists, return the user
	if err == sql.ErrNoRows {
		// Create user if doesn't exist
		username := strings.Split(UserDetail.Email, "@")[0] // basic username
		err = db.QueryRow(`
			INSERT INTO users (google_id, email, name, username, avatar_url)
			VALUES ($1, $2, $3, $4, $5)
			RETURNING id`, UserDetail.GoogleID, UserDetail.Email, UserDetail.Name, username, UserDetail.Avatar).
			Scan(&user.ID)

		if err != nil {
			log.Fatal("Error creating user:", err)
		}

		// Set user details
		user.GoogleID = UserDetail.GoogleID
		user.Email = UserDetail.Email
		user.Username = UserDetail.GivenName + "_" + UserDetail.FamilyName // basic username
		user.Avatar = UserDetail.Avatar
	}

	return user, err
}
