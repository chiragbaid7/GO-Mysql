package main

import (
	"log"
	"net/http"

	"github.com/chiragbaid7/GO-Mysql/pkg/routes"
	"github.com/gorilla/mux"
)

func main() {
	/*
		return new router instance that implements http.Handler interface
		registers routes and controller/handler fn
	*/
	router := mux.NewRouter()
	routes.BooksStoreRoutes(router)
	log.Fatal(http.ListenAndServe(":8080", router))
}
