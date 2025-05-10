package database

import (
	"bookhub/internal/types"
	"database/sql"
	"fmt"
	"log"

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

func GetUser(db *sql.DB, username string) (types.User, error) {
	query := `SELECT id, username, name, email, password FROM users WHERE username = $1;`
	var user types.User

	// return single row
	err := db.QueryRow(query, username).Scan(&user.ID, &user.Username, &user.Name, &user.Email, &user.Password)

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
