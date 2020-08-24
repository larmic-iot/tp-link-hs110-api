package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"

	"tp-link-hs110-api/api/client"
)

const (
	printDebug  = false
	timeoutInMs = 500
)

func Index(w http.ResponseWriter, _ *http.Request) {
	_, _ = fmt.Fprintln(w, "Hello tp-link-hs110-api!")
}

func InfoHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["ip"]

	socketClient := client.NewTpLinkHS110Client(key, timeoutInMs, printDebug)

	response, err := socketClient.RequestInfo()

	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		w.Header().Add("Content-Type", "text/plain")
		_, _ = w.Write([]byte(key + " not found."))
	}

	w.Header().Add("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(response)
}

func EnergyHandler(w http.ResponseWriter, r *http.Request) {
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
