package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/dlatyshev/GoRest/bookapi/models"
	"github.com/gorilla/mux"
)

func GetAllBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(models.GetAllBooks())
	if err != nil {
		log.Printf("Error encoding books to JSON: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func GetBookById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	id := params["id"]
	parsedId, err := strconv.Atoi(id)
	if err != nil {
		http.Error(w, "Invalid book ID", http.StatusBadRequest)
		return
	}
	book, found := models.FindBookById(parsedId)
	if !found {
		http.Error(w, "Book not found", http.StatusNotFound)
		return
	}
	err = json.NewEncoder(w).Encode(book)
	if err != nil {
		log.Printf("Error encoding book to JSON: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func CreateBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var book models.Book
	err := json.NewDecoder(r.Body).Decode(&book)
	if err != nil {
		http.Error(w, "Invalid book data", http.StatusBadRequest)
		return
	}
	newBookId := len(models.GetAllBooks()) + 1
	book.Id = newBookId
	models.AddBook(book)
	w.WriteHeader(http.StatusCreated)
	err = json.NewEncoder(w).Encode(book)
	if err != nil {
		log.Printf("Error encoding book to JSON: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func UpdateBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	id := params["id"]
	parsedId, err := strconv.Atoi(id)
	if err != nil {
		http.Error(w, "Invalid book ID", http.StatusBadRequest)
		return
	}
	book, found := models.FindBookById(parsedId)
	if !found {
		http.Error(w, "Book not found", http.StatusNotFound)
		return
	}
	err = json.NewDecoder(r.Body).Decode(&book)
	if err != nil {
		http.Error(w, "Invalid book data", http.StatusBadRequest)
		return
	}
	models.UpdateBook(parsedId, book)
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(book)
	if err != nil {
		log.Printf("Error encoding book to JSON: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func DeleteBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	id := params["id"]
	parsedId, err := strconv.Atoi(id)
	if err != nil {
		http.Error(w, "Invalid book ID", http.StatusBadRequest)
		return
	}
	models.DeleteBook(parsedId)
	w.WriteHeader(http.StatusOK)
}
