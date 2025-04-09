package pizzaapi

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strconv"
)

var (
	port string = "8080"
)

type Pizza struct {
	Id    int     `json:"id"`
	Name  string  `json:"name"`
	Size  string  `json:"size"`
	Price float64 `json:"price"`
}

// Database simulation
var db = []Pizza{
	{Id: 1, Name: "Margherita", Size: "Medium", Price: 8.99},
	{Id: 2, Name: "Pepperoni", Size: "Large", Price: 10.99},
	{Id: 3, Name: "Vegetarian", Size: "Small", Price: 7.99},
	{Id: 4, Name: "BBQ Chicken", Size: "Large", Price: 12.99},
	{Id: 5, Name: "Hawaiian", Size: "Medium", Price: 9.99},
}

func Run() {
	log.Println("Starting pizza API server on port", port)
	router := mux.NewRouter()
	router.HandleFunc("/pizzas", GetAllPizzas).Methods("GET")
	router.HandleFunc("/pizza/{id}", GetPizzaById).Methods("GET")
	log.Fatal(http.ListenAndServe(":"+port, router))
}

func FindPizzaById(id int) (Pizza, bool) {
	for _, pizza := range db {
		if pizza.Id == id {
			return pizza, true
		}
	}
	return Pizza{}, false
}

func GetPizzaById(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	params := mux.Vars(request)
	id := params["id"]
	pizzaId, err := strconv.Atoi(id)
	if err != nil {
		http.Error(writer, "Invalid pizza ID", http.StatusBadRequest)
		return
	}
	pizza, found := FindPizzaById(pizzaId)
	if !found {
		http.Error(writer, "Pizza not found", http.StatusNotFound)
		return
	}
	writer.WriteHeader(http.StatusOK)
	json.NewEncoder(writer).Encode(pizza)
}

func GetAllPizzas(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
	json.NewEncoder(writer).Encode(db)
}
