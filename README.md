# Bookhub is a library where users can create books and store them to the database

## Technology
- Golang
- PostgreSQL

### Installation
1. Clone this repository
2. Run the following to install the go files
```
go run tidy
```
3. Create a .env with the following database for psql
* SECRET_KEY=
* DB_PASSWORD=
* DB_USER=
* DB_NAME=
* GOOGLE_CLIENT_ID=
```
go run main.go
```

### Create Database in PostgreSQL
1. Create the database in PostgreSQL
```
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

  CREATE TABLE IF NOT EXISTS books (
		id INTEGER PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
		title TEXT NOT NULL,
		author TEXT NOT NULL,
		pages INT NOT NULL,
		publisher TEXT NOT NULL,
		isbn TEXT UNIQUE NOT NULL,
		description TEXT,
		published_at DATE,
		genre TEXT,
		cover_img TEXT DEFAULT ''
	);
```

### Starting the server
1. Run the main.go file
2. This should run on the port localhost:8080