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
3. Run the main.go file
```
go run main.go
```
4. This should run on the port localhost:8080