package controller

import (
	"net/http"
	"strconv"

	"github.com/chiragbaid7/GO-Mysql/pkg/utils"
	"github.com/gorilla/mux"
)

func GetBooks(r *http.Request, w http.ResponseWriter) {
	//get the books from database
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	//the data has to be json encoded
}

func GetBook(r *http.Request, w http.ResponseWriter) {
	vars := mux.Vars(r)
	book_id, _ := strconv.Atoi(vars["id"])
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusOK)
}

func deleteBook(r *http.Request, w http.ResponseWriter) {
	vars := mux.Vars(r)
	book_id, _ := strconv.Atoi(vars["id"])

}

func CreateBook(r *http.Request, w http.ResponseWriter) {
	//send to utils to decode the data
	utils.ParseBody()
}
