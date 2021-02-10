package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/Draiz25/Learning/Books-list/database"
	"github.com/gorilla/mux"
)

type Book struct {
	ID     int    `json:"id" xml:"id"`
	Title  string `json:"title" xml:"title"`
	Author string `json:"author" xml:"author"`
	Year   string `json:"year" xml:"year"`
}

var books []Book

// GetBooks send all the books
func GetBooks(w http.ResponseWriter, r *http.Request) {
	book := database.LoadDatabase()
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(book)
}

// GetBook send a single based on the id books
func GetBook(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)
	id := params["id"]
	book := database.LoadSingleRow(id)
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(book)
}

// AddBook adds a single book to the books
func AddBook(w http.ResponseWriter, r *http.Request) {
	var book database.Book
	json.NewDecoder(r.Body).Decode(&book)
	bookID := database.AddDetails(book)
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(bookID)
}

// UpdateBook upadate the information in a single book
func UpdateBook(w http.ResponseWriter, r *http.Request) {
	var book database.Book
	json.NewDecoder(r.Body).Decode(&book)
	rowsAffected := database.UpdateTables(book)
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(rowsAffected)

}

// RemoveBook removes a single book from the books
func RemoveBook(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	id := params["id"]
	rowsDeleted := database.RemoveDetails(id)
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(rowsDeleted)
}
