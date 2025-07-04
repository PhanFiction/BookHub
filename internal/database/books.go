package database

import (
	"bookhub/internal/types"
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

// Create a book table if it doesnt exist
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
		published_at DATE,
		genre TEXT,
		cover_img TEXT DEFAULT ''
	);
	`
	_, err := db.Exec(query)

	if err != nil {
		log.Fatal("Error creating books table:", err)
	}

	fmt.Println("Books table created or already exists.")
}

// Create saved book table if it doesnt exist
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
func FetchBooks(db *sql.DB, query string) []types.BookDetails {
	rows, err := db.Query(query) // Return all rows of books from the books table

	if err != nil {
		log.Fatal("Error fetching book:", err)
	}

	var title, author, publisher, isbn, description, publishedAt, genre, cover_img string
	var id int
	var pages int
	book := []types.BookDetails{}

	// Iterate through the rows and scan the values into variables
	for rows.Next() {
		err := rows.Scan(&id, &title, &author, &pages, &publisher, &isbn, &description, &publishedAt, &genre, &cover_img)
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
			Genre:       genre,
			CoverImg:    cover_img,
		})

		// fmt.Printf("Title: %s\nAuthor: %s\nPages: %d\nPublisher: %s\nISBN: %s\nDescription: %s\nPublished At: %s\n", title, author, pages, publisher, isbn, description, publishedAt)
	}

	return book
}

// Fetch a single book from the database
func FetchSingleBook(db *sql.DB, bookId string) types.BookDetails {
	var title, author, publisher, isbn, description, publishedAt, genre, cover_img string
	var id int
	var pages int

	query := `SELECT * FROM books	WHERE id = $1;`

	err := db.QueryRow(query, bookId).Scan(&id, &title, &author, &pages, &publisher, &isbn, &description, &publishedAt, &genre, &cover_img)

	if err != nil {
		log.Fatal("Error fetching single book:", err)
	}

	return types.BookDetails{
		ID:          id,
		Title:       title,
		Author:      author,
		Pages:       pages,
		Publisher:   publisher,
		ISBN:        isbn,
		Description: description,
		PublishedAt: publishedAt,
		Genre:       genre,
		CoverImg:    cover_img,
	}
}

// Create a book in the database
func CreateBook(db *sql.DB, BookDetails types.BookDetails) {
	query := `
		INSERT INTO books (title, author, pages, publisher, isbn, description, published_at, genre, cover_img)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)
		ON CONFLICT (isbn) DO NOTHING;
	`
	fmt.Println("Creating book:", BookDetails.Title)
	_, err := db.Exec(query, BookDetails.Title, BookDetails.Author, BookDetails.Pages, BookDetails.Publisher, BookDetails.ISBN, BookDetails.Description, BookDetails.PublishedAt, BookDetails.Genre, BookDetails.CoverImg)

	if err != nil {
		log.Fatal("Error creating book:", err)
	}

	fmt.Println("Book created in database.")
}

// Update a book in the database
func UpdateBook(db *sql.DB, BookDetails types.BookDetails, bookId string) error {
	query := `
	UPDATE books
	SET title = $1, author = $2, pages = $3, publisher = $4, isbn = $5, description = $6, published_at = $7, genre = $8, cover_img = $9
	WHERE id = $10;
	`

	fmt.Println(BookDetails.Genre)

	_, err := db.Exec(query, BookDetails.Title, BookDetails.Author, BookDetails.Pages, BookDetails.Publisher, BookDetails.ISBN, BookDetails.Description, BookDetails.PublishedAt, BookDetails.Genre, BookDetails.CoverImg, bookId)

	if err != nil {
		log.Fatal("Error updating book:", err)
	}

	fmt.Println("Book updated in database.")

	return err
}

// Delete a book from the database
func DeleteBook(db *sql.DB, bookID string) error {
	query := `
	DELETE FROM books
	WHERE id = $1;
	`
	_, err := db.Exec(query, bookID)

	if err != nil {
		log.Fatal("Error deleting book:", err)
	}

	fmt.Println("Book deleted from database.")

	return err
}

// Save the book to the saved_books table
// This function saves a book for a user, preventing duplicates
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
