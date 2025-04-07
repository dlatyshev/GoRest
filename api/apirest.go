package api

import (
	"fmt"
	"log"
	"net/http"
)

// GetGreet is a simple HTTP handler that responds with a greeting message.
func GetGreet(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hi, I am web server")
}

func RequestHandler() {
	http.HandleFunc("/", GetGreet)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
