package handlers

import (
	"bookhub/internal/database"
	"bookhub/internal/session"
	"bookhub/internal/types"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// Return single book from the database
func FetchBooksHandler(w http.ResponseWriter, r *http.Request) {
	data := database.FetchBooks(database.DB, "SELECT * FROM books;")
	json.NewEncoder(w).Encode(data)
}

// Fetch single book from the database
func FetchSingleBookHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r) // extract route variables
	bookId := vars["id"]
	//bookId_to_int, _ := strconv.Atoi(bookId)
	data := database.FetchSingleBook(database.DB, bookId)
	json.NewEncoder(w).Encode(data)
}

func UpdateBookHandler(w http.ResponseWriter, r *http.Request) {
	var body types.BookDetails
	json.NewDecoder(r.Body).Decode(&body)
	vars := mux.Vars(r) // extract route variables
	bookId := vars["id"]

	BookData := types.BookDetails{
		Title:       body.Title,
		Author:      body.Author,
		Pages:       body.Pages,
		Publisher:   body.Publisher,
		ISBN:        body.ISBN,
		Description: body.Description,
		PublishedAt: body.PublishedAt,
		Genre:       body.Genre,
	}

	err := database.UpdateBook(database.DB, BookData, bookId)

	if err != nil {
		http.Error(w, "Error updating book", http.StatusInternalServerError)
		return
	}

	data := types.Data{
		Message: "Book updated successfully",
	}

	json.NewEncoder(w).Encode(data)
}

func CreateBookHandler(w http.ResponseWriter, r *http.Request) {
	var body types.BookDetails
	json.NewDecoder(r.Body).Decode(&body)

	BookData := types.BookDetails{
		Title:       body.Title,
		Author:      body.Author,
		Pages:       body.Pages,
		Publisher:   body.Publisher,
		ISBN:        body.ISBN,
		Description: body.Description,
		PublishedAt: body.PublishedAt,
		Genre:       body.Genre,
	}

	fmt.Println(body)

	database.CreateBook(database.DB, BookData)
	data := types.Data{
		Message: "Book created successfully",
	}
	json.NewEncoder(w).Encode(data)
}

func DeleteBookHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r) // extract route variables
	bookId := vars["id"]
	err := database.DeleteBook(database.DB, bookId)

	if err != nil {
		http.Error(w, "Error updating book", http.StatusInternalServerError)
		return
	}

	data := types.Data{
		Message: "Book successfully deleted",
	}

	json.NewEncoder(w).Encode(data)
}

// SaveBookHandler saves a book for the user
func SaveBookHandler(w http.ResponseWriter, r *http.Request) {
	// save book and delete saved book for user
	bookIdStr := mux.Vars(r)["id"] // extract bookid from params

	// Get user ID from session
	session, _ := session.Store.Get(r, "session")
	userID, _ := session.Values["user_id"].(int)

	// Convert bookId from string to int
	bookId, err := strconv.Atoi(bookIdStr)
	if err != nil {
		http.Error(w, "Invalid book ID", http.StatusBadRequest)
		return
	}

	// Check if the book is already saved
	exists := database.CheckIfBookExists(database.DB, strconv.Itoa(userID), bookId)
	fmt.Println("Book exists:", exists)
	if exists {
		// If the book is already saved, delete it
		err = database.DeleteSavedBook(database.DB, userID, bookId)
		if err != nil {
			http.Error(w, "Error deleting saved book", http.StatusInternalServerError)
			return
		}
		data := types.Data{
			Message: "Book deleted successfully",
		}
		json.NewEncoder(w).Encode(data)
		return
	}

	database.SaveBook(database.DB, userID, bookId)
	data := types.Data{
		Message: "Book saved successfully",
	}
	json.NewEncoder(w).Encode(data)
}
