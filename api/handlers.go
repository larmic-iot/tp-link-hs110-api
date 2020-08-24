package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"

	"tp-link-hs110-api/api/client"
	"tp-link-hs110-api/api/model"
)

const (
	printDebug  = false
	timeoutInMs = 500
)

func InfoHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["ip"]

	socketClient := client.NewTpLinkHS110Client(key, timeoutInMs, printDebug)

	response, err := socketClient.RequestInfo()

	if err != nil {
		w.Header().Add("Content-Type", "text/plain; charset=UTF-8")
		w.WriteHeader(http.StatusNotFound)
		_, _ = fmt.Fprintf(w, "%s not found!", key)
	}

	ledState := model.On

	if response.LedOff == 1 {
		ledState = model.Off
	}

	powerState := model.On

	if response.RelayState == 0 {
		powerState = model.Off
	}

	info := model.Info{
		Ip:              key,
		Port:            9999,
		Name:            response.Alias,
		Icon:            response.Icon,
		Model:           response.Model,
		MacAddress:      response.MacAddress,
		SoftwareVersion: response.SoftwareVersion,
		HardwareVersion: response.HardwareVersion,
		Led:             ledState,
		Power:           powerState,
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

func OpenApiDocumentation(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/yaml; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	dat, _ := ioutil.ReadFile("open-api-3.yaml")
	_, _ = w.Write(dat)
}
