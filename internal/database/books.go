package database

import (
	"bookhub/internal/types"
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

func CreateBookTable(db *sql.DB) {
	query := `
	CREATE TABLE IF NOT EXISTS books (
		id INTEGER PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
		title TEXT NOT NULL,
		author TEXT NOT NULL,
		pages INT NOT NULL,
		publisher TEXT NOT NULL,
		isbn TEXT UNIQUE NOT NULL,
		description TEXT,
		published_at DATE
	);
	`
	_, err := db.Exec(query)

	if err != nil {
		log.Fatal("Error creating books table:", err)
	}

	fmt.Println("Books table created or already exists.")
}

func CreateSavedBooksTable(db *sql.DB) {
	query := `
	CREATE TABLE IF NOT EXISTS saved_books (
	user_id INT REFERENCES users(id) ON DELETE CASCADE,
	book_id INT REFERENCES books(id) ON DELETE CASCADE,
	saved_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
	PRIMARY KEY (user_id, book_id)
);
	`
	_, err := db.Exec(query)

	if err != nil {
		log.Fatal("Error creating saved_books table:", err)
	}

	fmt.Println("Saved books table created or already exists.")
}

// Fetch book from db based on data provided
func FetchBook(db *sql.DB, query string) []types.BookDetails {
	rows, err := db.Query(query) // Return all rows of books from the books table

	if err != nil {
		log.Fatal("Error fetching book:", err)
	}

	var title, author, publisher, isbn, description, publishedAt string
	var id int
	var pages int
	book := []types.BookDetails{}

	for rows.Next() {
		err := rows.Scan(&id, &title, &author, &pages, &publisher, &isbn, &description, &publishedAt)
		if err != nil {
			log.Fatal("Error scanning book:", err)
		}

		book = append(book, types.BookDetails{
			ID:          id,
			Title:       title,
			Author:      author,
			Pages:       pages,
			Publisher:   publisher,
			ISBN:        isbn,
			Description: description,
			PublishedAt: publishedAt,
		})

		// fmt.Printf("Title: %s\nAuthor: %s\nPages: %d\nPublisher: %s\nISBN: %s\nDescription: %s\nPublished At: %s\n", title, author, pages, publisher, isbn, description, publishedAt)
	}

	return book
}

func CreateBook(db *sql.DB, BookDetails types.BookDetails) {
	query := `
		INSERT INTO books (title, author, pages, publisher, isbn, description, published_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7)
		ON CONFLICT (isbn) DO NOTHING;
	`
	fmt.Println("Creating book:", BookDetails.Title)
	_, err := db.Exec(query, BookDetails.Title, BookDetails.Author, BookDetails.Pages, BookDetails.Publisher, BookDetails.ISBN, BookDetails.Description, BookDetails.PublishedAt)

	if err != nil {
		log.Fatal("Error creating book:", err)
	}

	fmt.Println("Book created in database.")
}

func UpdateBook(db *sql.DB, bookID int, title, author string, pages int, publisher, isbn, description, publishedAt string) {
	query := `
	UPDATE books
	SET title = $1, author = $2, pages = $3, publisher = $4, isbn = $5, description = $6, published_at = $7
	WHERE id = $8;
	`
	_, err := db.Exec(query, title, author, pages, publisher, isbn, description, publishedAt, bookID)

	if err != nil {
		log.Fatal("Error updating book:", err)
	}

	fmt.Println("Book updated in database.")
}

func DeleteBook(db *sql.DB, bookID int) {
	query := `
	DELETE FROM books
	WHERE id = $1;
	`
	_, err := db.Exec(query, bookID)

	if err != nil {
		log.Fatal("Error deleting book:", err)
	}

	fmt.Println("Book deleted from database.")
}

func SaveBook(db *sql.DB, userID, bookID int) {
	query := `
	INSERT INTO saved_books (user_id, book_id)
	VALUES ($1, $2)
	ON CONFLICT DO NOTHING;
	`
	_, err := db.Exec(query, userID, bookID)

	if err != nil {
		log.Fatal("Error saving book:", err)
	}

	fmt.Println("Book saved to database.")
}
