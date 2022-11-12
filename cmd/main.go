package main

import (
	"fmt"
	"log"
	"net/http"

	_ "github.com/chiragbaid7/GO-Mysql/pkg/models"
	// "github.com/chiragbaid7/GO-Mysql/pkg/routes"
	"github.com/gorilla/mux"
)

func main() {
	/*
		return new router instance that implements http.Handler interface
		registers routes and controller/handler fn
	*/
	fmt.Println("chirag")
	router := mux.NewRouter()
	// routes.BooksStoreRoutes(router)
	log.Fatal(http.ListenAndServe(":8080", router))
}
