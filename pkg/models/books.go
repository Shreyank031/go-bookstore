package models

import (
	"github.com/Shreyank031/go-bookstore/pkg/config"
	"github.com/jinzhu/gorm"
)

var db *gorm.DB //db holds the databse connection instance

type Book struct {
	gorm.Model

	//gorm for database mapping and json for JSON serialization/deserialization
	Name        string `gorm:""json:"name"`
	Author      string `json:"author"`
	Publication string `json:"publication"`
}

func init() {
	config.Connect()
	db = config.GetDB()
	//db.AutoMigrate(&Book{})-> Create the Book table in the database (if it doesn't exist already)
	db.AutoMigrate(&Book{})
}

// The CreateBook function needs to modify the Book struct instance, and in Go, to modify a struct, you need to use a pointer receiver.
func (d *Book) CreateBook() *Book {
	db.NewRecord(d)
	db.Create(&d)

	return d
}

func GetAllBooks() []Book {
	var Books []Book
	db.Find(&Books)
	return Books
}

func GetBookById(id int64) (*Book, *gorm.DB) {
	var getBook Book
	db := db.Where("ID=?", id).Find(&getBook)
	return &getBook, db //return the query result db and memory address of getBook which is pointer to the book struct containing the retreived book data
}

func DeleteBookById(ID int64) *Book {
	var book Book
	db.Where("ID=?", ID).Delete(book)
	return &book
}
