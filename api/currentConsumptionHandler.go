package api

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/gorilla/mux"

	"tp-link-hs110-api/api/client"
	clientModel "tp-link-hs110-api/api/client/model"
	"tp-link-hs110-api/api/model"
)

func CurrentConsumptionHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	ip := vars["ip"]
	year, month, day := time.Now().Date()

	socketClient := client.NewTpLinkHS110Client(ip, timeoutInMs, printDebug)

	currentEnergyResponse, err := socketClient.RequestCurrentEnergyStatistics()
	dailyEnergyResponse, err2 := socketClient.RequestDailyEnergyStatistics(year, month, day)

	if err != nil || err2 != nil {
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

	energy := mapEnergyModel(currentEnergyResponse, dailyEnergyResponse)

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(energy)
}

func mapEnergyModel(eMeterInfo clientModel.EMeterInfo, dailyEMeterInfo clientModel.DayStatEMeterInfo) model.Consumption {
	return model.Consumption{
		CurrentMW: eMeterInfo.Power,
		TodayWH:   dailyEMeterInfo.EnergyWattHours,
		TotalWH:   eMeterInfo.TotalWattHours,
	}
}
