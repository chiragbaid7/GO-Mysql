package routes

import (
	"github.com/chiragbaid7/GO-Mysql/pkg/controller"
	"github.com/gorilla/mux"
)

func BooksStoreRoutes(router *mux.Router) {
	router.HandleFunc("/", controller.GetMovies).Methods("GET")
	router.HandleFunc("/{id}", controller.GetMovie).Methods("GET")
	router.HandleFunc("/", controller.CreateBook).Methods("POST")
	router.HandleFunc("/", controller.DeleteBook).Methods("DELETE")
}
