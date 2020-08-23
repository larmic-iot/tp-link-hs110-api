package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"

	"tp-link-hs110-api/client"
)

const (
	printDebug  = false
	timeoutInMs = 500
)

func main() {
	log.Println("Hello tp-link-hs110-api!")

	router := mux.NewRouter().StrictSlash(true)
	router.NotFoundHandler = Handle404()

	router.HandleFunc("/api/{ip}", infoHandler).Methods("GET")
	router.HandleFunc("/api/{ip}/energy", energyHandler).Methods("GET")

	log.Fatal(http.ListenAndServe(":8080", router))
}

func infoHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["ip"]

	socketClient := client.NewTpLinkHS110Client(key, timeoutInMs, printDebug)

	response, err := socketClient.RequestInfo()

	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		w.Header().Add("Content-Type", "text/plain")
		_, _ = w.Write([]byte(key + " not found."))
	}

	marshal, err := json.Marshal(response)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Header().Add("Content-Type", "text/plain")
		_, _ = w.Write([]byte(key + " throws internal server error."))
	}

	w.Header().Add("Content-Type", "application/json")
	_, _ = w.Write(marshal)
}

func energyHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["ip"]

	socketClient := client.NewTpLinkHS110Client(key, timeoutInMs, printDebug)

	response, err := socketClient.RequestCurrentEnergyStatistics()

	if err == nil {
		_, _ = w.Write([]byte(response))
	} else {
		w.WriteHeader(http.StatusNotFound)
		_, _ = w.Write([]byte(key + " not found."))
	}
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
