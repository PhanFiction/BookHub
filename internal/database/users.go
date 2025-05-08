package database

import (
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

func GetUsersTable(db *sql.DB) []string {
	query := `SELECT username FROM usernames;`

	rows, _ := db.Query(query)

	var usernames []string

	for rows.Next() {
		var username string
		rows.Scan(&username)
		usernames = append(usernames, username)
	}

	fmt.Println(usernames)
	return usernames
}

func SaveUser(db *sql.DB, email, name, username, password string) {
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
