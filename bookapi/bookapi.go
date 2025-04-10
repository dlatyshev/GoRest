package bookapi

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/dlatyshev/GoRest/bookapi/utils"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

const (
	apiPrefix           string = "/api/v1"
	bookResourcePrefix  string = apiPrefix + "/book"
	booksResourcePrefix string = apiPrefix + "/books"
)

var (
	port string
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	port = os.Getenv("port")
}

func StartServer() {
	fmt.Println("Starting server on port", port)
	router := mux.NewRouter()
	utils.BuildBookResource(router, bookResourcePrefix)
	utils.BuildBooksResource(router, booksResourcePrefix)
	log.Fatal(http.ListenAndServe(":"+port, router))
}
