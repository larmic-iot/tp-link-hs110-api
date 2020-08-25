package api

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"

	"tp-link-hs110-api/api/client"
)

func EnergyHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["ip"]

	socketClient := client.NewTpLinkHS110Client(key, timeoutInMs, printDebug)

	response, err := socketClient.RequestCurrentEnergyStatistics()

	if err == nil {
		_ = json.NewEncoder(w).Encode(response)
	} else {
		w.WriteHeader(http.StatusNotFound)
		_, _ = w.Write([]byte(key + " not found."))
	}
}
