package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"

	"tp-link-hs110-api/client"
)

const (
	printDebug = false
)

func main() {
	log.Println("Hello tp-link-hs110-api!")

	//fridgeClient := client.NewTpLinkHS110Client("10.0.0.78", printDebug)

	//log.Printf("Request info: %s", fridgeClient.RequestInfo())
	//log.Printf("Request energy: %s", fridgeClient.RequestCurrentEnergyStatistics())
	//log.Printf("Request energy: %s", fridgeClient.RequestSwitchOff())
	//log.Printf("Request energy: %s", fridgeClient.RequestSwitchOn())

	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/api/hello/{ip}", testHandler).Methods("GET")
	log.Fatal(http.ListenAndServe(":8080", router))
}

func testHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["ip"]

	socketClient := client.NewTpLinkHS110Client(key, printDebug)

	w.Write([]byte(socketClient.RequestInfo()))
}
