package routes

import (
	"github.com/chiragbaid7/GO-Mysql/pkg/controller"
	"github.com/gorilla/mux"
)

func BooksStoreRoutes(router *mux.Router) {
	router.HandleFunc("/", controller.GetBooks).Methods("GET")
	router.HandleFunc("/{id}", controller.GetBook).Methods("GET")
	router.HandleFunc("/", controller.CreateBook).Methods("POST")
	router.HandleFunc("/{id}", controller.DeleteBook).Methods("DELETE")
}
