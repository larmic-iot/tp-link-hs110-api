package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"

	"tp-link-hs110-api/api"
)

func main() {
	log.Println("Hello tp-link-hs110-api!")

	router := mux.NewRouter().StrictSlash(true)
	router.NotFoundHandler = Handle404()

	router.HandleFunc("/", api.Index).Methods("GET")
	router.HandleFunc("/api/{ip}", api.InfoHandler).Methods("GET")
	router.HandleFunc("/api/{ip}/energy", api.EnergyHandler).Methods("GET")

	log.Fatal(http.ListenAndServe(":8080", router))
}

type ProtocolError struct {
	Code    int
	Message string
}

func Handle404() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		json.
			NewEncoder(w).
			Encode(
				ProtocolError{
					Code:    http.StatusNotFound,
					Message: http.StatusText(http.StatusNotFound),
				})
	})
}
