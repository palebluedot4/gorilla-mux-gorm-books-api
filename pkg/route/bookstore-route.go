package route

import (
	"github.com/gorilla/mux"

	"gorilla-mux-gorm-books-api/pkg/controller"
)

var RegisterBookstoreRoute = func(router *mux.Router) {
	router.HandleFunc("/books", controller.GetBooksHandler).Methods("GET")
	router.HandleFunc("/books/{isbn}", controller.GetBookByISBNHandler).Methods("GET")
	router.HandleFunc("/books", controller.CreateBookHandler).Methods("POST")
	router.HandleFunc("/books/{isbn}", controller.DeleteBookHandler).Methods("DELETE")
	router.HandleFunc("/boos/{isbn}", controller.UpdateBookHandler).Methods("PUT")
}
