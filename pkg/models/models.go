package models

import (
	"fmt"

	"github.com/chiragbaid7/GO-Mysql/pkg/configs"
	"gorm.io/gorm"
)

// Define your own Model -> table in your database.
type Book struct {
	gorm.Model
	Name   string
	Author string
	Book   string
}

var db *gorm.DB

func init() {
	fmt.Println("inside models")
	db = configs.Connect()
	fmt.Println("Connected to database")
	// like sync->auto migrate
	db.AutoMigrate(&Book{})
}

func GetBooksModel() {

}
