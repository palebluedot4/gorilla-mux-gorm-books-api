package model

import (
	"log"

	"gorm.io/gorm"

	"gorilla-mux-gorm-books-api/pkg/config"
)

var db *gorm.DB

type Book struct {
	gorm.Model
	ISBN   string  `gorm:"type:VARCHAR(100);NOT NULL"`
	Title  string  `gorm:"type:VARCHAR(100);NOT NULL" json:"title"`
	Author string  `gorm:"type:VARCHAR(100);NOT NULL" json:"author"`
	Price  float64 `gorm:"type:DOUBLE(11,2);NOT NULL" json:"price"`
	Stock  int64   `gorm:"type:INT;NOT NULL" json:"stock"`
}

func init() {
	config.ConnectDB()
	db = config.GetDB()
	if err := db.AutoMigrate(&Book{}); err != nil {
		log.Fatalf("error migrating database: %v", err)
	}
}

func (b *Book) CreateBook() *Book {
	db.Create(&b)
	return b
}

func GetBooks() []Book {
	var books []Book
	db.Find(&books)
	return books
}

func GetBookByISBN(ISBN int64) (*Book, *gorm.DB) {
	var books Book
	db := db.Where("ISBN = ?", ISBN).Find(&books)
	return &books, db
}

func DeleteBook(ISBN int64) *Book {
	var book Book
	db.Where("ISBN = ?", ISBN).Delete(&book)
	return &book
}
