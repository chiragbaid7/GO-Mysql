package utils

import (
	"encoding/json"
	"net/http"
)

type Book struct {
	Name   string `json:"name"`
	Author string `json:"author"`
	Book   string `json:"book"`
}

func ParseBody(r *http.Request) *Book {
	//create a decoding pipeline
	var book *Book
	json.NewDecoder(r.Body).Decode(&book)
	return book
}
