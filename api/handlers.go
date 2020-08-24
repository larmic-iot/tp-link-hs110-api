package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"

	"tp-link-hs110-api/api/client"
	"tp-link-hs110-api/api/model"
)

const (
	printDebug  = false
	timeoutInMs = 500
)

func Index(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(http.StatusOK)
	_, _ = fmt.Fprintln(w, "Hello tp-link-hs110-api!")
}

func InfoHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["ip"]

	socketClient := client.NewTpLinkHS110Client(key, timeoutInMs, printDebug)

	response, err := socketClient.RequestInfo()

	if err != nil {
		w.Header().Add("Content-Type", "text/plain")
		w.WriteHeader(http.StatusNotFound)
		_, _ = fmt.Fprintf(w, "%s not found!", key)
	}

	led := false

	if response.LedOff == 0 {
		led = true
	}

	info := model.Info{
		Identifier:      key,
		Name:            response.Alias,
		Icon:            response.Icon,
		Model:           response.Model,
		MacAddress:      response.MacAddress,
		SoftwareVersion: response.SoftwareVersion,
		HardwareVersion: response.HardwareVersion,
		Led:             led,
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(info)
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
