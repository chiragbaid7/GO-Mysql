package main

import (
	"log"
	"net/http"

	"githbub.com/chiragbaid7/Go-Mysql/pkg/routes"
	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	routes.BooksStoreRoutes(router)
	log.Fatal(http.ListenAndServe(":8080", router))
}
