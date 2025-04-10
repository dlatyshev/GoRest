package utils

import (
	"github.com/dlatyshev/GoRest/bookapi/handlers"
	"github.com/gorilla/mux"
)

func BuildBookResource(router *mux.Router, prefix string) {
	router.HandleFunc(prefix+"/{id}", handlers.GetBookById).Methods("GET")
	router.HandleFunc(prefix, handlers.CreateBook).Methods("POST")
	router.HandleFunc(prefix+"/{id}", handlers.UpdateBook).Methods("PUT")
	router.HandleFunc(prefix+"/{id}", handlers.DeleteBook).Methods("DELETE")
}

func BuildBooksResource(router *mux.Router, prefix string) {
	router.HandleFunc(prefix, handlers.GetAllBooks).Methods("GET")
}
