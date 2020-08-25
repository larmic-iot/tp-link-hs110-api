package api

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"

	"tp-link-hs110-api/api/client"
	clientModel "tp-link-hs110-api/api/client/model"
	"tp-link-hs110-api/api/model"
)

func EnergyHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	ip := vars["ip"]

	socketClient := client.NewTpLinkHS110Client(ip, timeoutInMs, printDebug)

	response, err := socketClient.RequestCurrentEnergyStatistics()

	if err != nil {
		w.Header().Add("Content-Type", "text/plain; charset=UTF-8")
		w.WriteHeader(http.StatusNotFound)

		_ = json.
			NewEncoder(w).
			Encode(
				model.ProtocolError{
					Code:    http.StatusNotFound,
					Message: ip + " not found!",
				})
		return
	}

	energy := mapEnergyModel(response)

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(energy)
}

func mapEnergyModel(eMeterInfo clientModel.EMeterInfo) model.Energy {
	return model.Energy{
		Watt: eMeterInfo.Power,
	}
}
