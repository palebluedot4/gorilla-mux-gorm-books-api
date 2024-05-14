package controller

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"

	"gorilla-mux-gorm-books-api/pkg/model"
	"gorilla-mux-gorm-books-api/pkg/utils"
)

var Book model.Book

func GetBooksHandler(w http.ResponseWriter, r *http.Request) {
	books := model.GetBooks()

	booksJSON, _ := json.Marshal(books)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	if _, err := w.Write(booksJSON); err != nil {
		log.Printf("error writing response: %v", err)
	}
}

func GetBookByISBNHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	isbn := vars["isbn"]
	bookISBN, err := strconv.ParseInt(isbn, 10, 64)
	if err != nil {
		log.Printf("err: %v", err)
	}

	book, _ := model.GetBookByISBN(bookISBN)
	bookJSON, _ := json.Marshal(book)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	if _, err := w.Write(bookJSON); err != nil {
		log.Printf("error writing response: %v", err)
	}
}

func CreateBookHandler(w http.ResponseWriter, r *http.Request) {
	newBook := &model.Book{}
	utils.ParseBody(r, newBook)
	book := newBook.CreateBook()
	bookJSON, _ := json.Marshal(book)
	w.WriteHeader(http.StatusOK)

	if _, err := w.Write(bookJSON); err != nil {
		log.Printf("error writing response: %v", err)
	}

}

func DeleteBookHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	isbn := vars["isbn"]
	bookISBN, err := strconv.ParseInt(isbn, 10, 64)
	if err != nil {
		log.Printf("error parsing ISBN: %v", err)
		http.Error(w, "invalid ISBN", http.StatusBadRequest)
		return
	}

	model.DeleteBook(bookISBN)

	w.WriteHeader(http.StatusOK)
	if _, err := w.Write([]byte("book deleted successfully")); err != nil {
		log.Printf("error writing response: %v", err)
	}
}

func UpdateBookHandler(w http.ResponseWriter, r *http.Request) {
	updateBook := &model.Book{}
	utils.ParseBody(r, updateBook)
	vars := mux.Vars(r)
	isbn := vars["isbn"]
	bookISBN, err := strconv.ParseInt(isbn, 10, 64)
	if err != nil {
		log.Printf("error parsing ISBN: %v", err)
		http.Error(w, "invalid ISBN", http.StatusBadRequest)
		return
	}

	book, db := model.GetBookByISBN(bookISBN)
	if book == nil {
		http.Error(w, "book not found", http.StatusNotFound)
		return
	}
	if updateBook.ISBN != "" {
		book.ISBN = updateBook.ISBN
	}
	if updateBook.Title != "" {
		book.Title = updateBook.Title
	}
	if updateBook.Author != "" {
		book.Author = updateBook.Author
	}
	if updateBook.Price != 0 {
		book.Price = updateBook.Price
	}
	if updateBook.Stock != 0 {
		book.Stock = updateBook.Stock
	}

	db.Save(book)
	bookJSON, _ := json.Marshal(book)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	if _, err := w.Write(bookJSON); err != nil {
		log.Printf("error writing response: %v", err)
	}
}
