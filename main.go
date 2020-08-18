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

	//fridgeClient := client.NewTpLinkHS110Client("10.0.0.78", printDebug)

	//log.Printf("Request info: %s", fridgeClient.RequestInfo())
	//log.Printf("Request energy: %s", fridgeClient.RequestCurrentEnergyStatistics())
	//log.Printf("Request energy: %s", fridgeClient.RequestSwitchOff())
	//log.Printf("Request energy: %s", fridgeClient.RequestSwitchOn())

	router := mux.NewRouter().StrictSlash(true)
	router.NotFoundHandler = Handle404()

	router.HandleFunc("/api/hello/{ip}", testHandler).Methods("GET")
	log.Fatal(http.ListenAndServe(":8080", router))
}

func testHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["ip"]

	socketClient := client.NewTpLinkHS110Client(key, timeoutInMs, printDebug)

	response, err := socketClient.RequestInfo()

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
