package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"time"

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
	newAuthorId := int(time.Now().UnixNano())
	book.Id = newBookId
	book.Author.Id = newAuthorId
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

	_, found := models.FindBookById(parsedId)
	if !found {
		http.Error(w, "Book not found", http.StatusNotFound)
		return
	}

	var updatedBook models.Book
	if err := json.NewDecoder(r.Body).Decode(&updatedBook); err != nil {
		http.Error(w, "Invalid book data", http.StatusBadRequest)
		return
	}

	if !models.UpdateBook(parsedId, updatedBook) {
		http.Error(w, "Book not found", http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
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

	_, found := models.FindBookById(parsedId)
	if !found {
		http.Error(w, "Book not found", http.StatusNotFound)
		return
	}

	models.DeleteBook(parsedId)
	w.WriteHeader(http.StatusOK)
	_, err = w.Write([]byte("Book with id " + id + " deleted successfully"))
	if err != nil {
		log.Printf("Error writing response: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
