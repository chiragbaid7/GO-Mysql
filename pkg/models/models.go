package models

import (
	"fmt"

	"github.com/chiragbaid7/GO-Mysql/pkg/configs"
	"gorm.io/gorm"
)

var db *gorm.DB

// Define your own Model -> table in your database.
type Book struct {
	gorm.Model
	Name   string `json:"name"`
	Author string `json:"author"`
	Book   string
}

func init() {
	db = configs.Connect()
	// like sync->auto migrate
	db.AutoMigrate(&Book{})
}

func CreateBook(book *Book) uint {
	db.Create(&book)
	fmt.Println("New Record created", book)
	return book.ID
}

func GetBooksModel() []Book {
	var books []Book
	db.Find(&books)
	return books
}

func GetBookModel(id int) Book {
	var book Book
	db.Find(&book, id)
	return book
}

func DeleteBookModel(id int) int {
	fmt.Println(id)
	db.Delete(&Book{}, id)
	return id
}
