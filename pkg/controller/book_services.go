package controller

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/chiragbaid7/GO-Mysql/pkg/models"
	"github.com/gorilla/mux"
)

func CreateBook(w http.ResponseWriter, r *http.Request) {
	//send to utils to decode the data
	book := &models.Book{}
	json.NewDecoder(r.Body).Decode(&book)
	models.CreateBook(book)
	w.Header().Set("Content-type", "application/json")
	json.NewEncoder(w).Encode(book)
	w.WriteHeader(http.StatusOK)
}

func GetBooks(w http.ResponseWriter, r *http.Request) {
	//get the books from database
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	books := models.GetBooksModel()
	json.NewEncoder(w).Encode(books)
	//or marshal(encode) and then write to w
	//res:=json.marshal(bookd)
	//w.write(res)
}

func GetBook(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	book_id, _ := strconv.Atoi(vars["id"])
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	book := models.GetBookModel(book_id)
	json.NewEncoder(w).Encode(book)
}

func DeleteBook(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	book_id, _ := strconv.Atoi(vars["id"])
	models.DeleteBookModel(book_id)
}
