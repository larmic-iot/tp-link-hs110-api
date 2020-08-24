package main

import (
	"log"
	"net/http"

	"tp-link-hs110-api/api"
)

func main() {
	log.Println("Hello tp-link-hs110-api!")

	router := api.NewRouter()

	log.Fatal(http.ListenAndServe(":8080", router))
}
