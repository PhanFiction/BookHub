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
		id SERIAL PRIMARY KEY,
		google_id TEXT UNIQUE NOT NULL,
		email TEXT UNIQUE NOT NULL,
		name TEXT,
		username TEXT,
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
