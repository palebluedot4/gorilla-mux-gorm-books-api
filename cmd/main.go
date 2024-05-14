package main

import (
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"

	"gorilla-mux-gorm-books-api/pkg/route"
)

func main() {
	router := mux.NewRouter()
	route.RegisterBookstoreRoute(router)
	http.Handle("/", router)
	log.Println("Listening and serving HTTP on localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
