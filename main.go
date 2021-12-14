package main

import (
	"log"
	"net/http"

	"github.com/gorilla/handlers"

	"tp-link-hs110-api/api"
)

func main() {
	log.Println("Hello tp-link-hs110-api! TEST")

	router := api.NewRouter()

	log.Fatal(http.ListenAndServe(":8080", handlers.CORS(handlers.AllowedOrigins([]string{"*"}))(router)))
}
