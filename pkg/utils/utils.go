package utils

import (
	"net/http"
)

type Book struct {
	Name   string `json:"name"`
	Author string `json:"author"`
	Book   string `json:"book"`
}

func ParseBody(r *http.Request, book *Book) {
	//create a decoding pipeline
	//decode expects a pointer
}
