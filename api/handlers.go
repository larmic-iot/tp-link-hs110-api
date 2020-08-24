package api

import (
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"

	"tp-link-hs110-api/api/client"
)

const (
	printDebug  = false
	timeoutInMs = 500
)

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
