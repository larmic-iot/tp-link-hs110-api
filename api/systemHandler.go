package api

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"

	"tp-link-hs110-api/api/client"
	clientModel "tp-link-hs110-api/api/client/model"
	"tp-link-hs110-api/api/model"
)

func GetSystemHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	ip := vars["ip"]

	socketClient := client.NewTpLinkHS110Client(ip, timeoutInMs, printDebug)

	response, err := socketClient.RequestInfo()

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

	system := mapClientModelToApiModel(ip, response)

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(system)
}

func mapClientModelToApiModel(ip string, systemInfo clientModel.SystemInfo) model.System {
	ledState := model.On

	if systemInfo.LedOff == 1 {
		ledState = model.Off
	}

	powerState := model.On

	if systemInfo.RelayState == 0 {
		powerState = model.Off
	}

	system := model.System{
		Ip:              ip,
		Port:            9999,
		Name:            systemInfo.Alias,
		Icon:            systemInfo.Icon,
		Model:           systemInfo.Model,
		MacAddress:      systemInfo.MacAddress,
		SoftwareVersion: systemInfo.SoftwareVersion,
		HardwareVersion: systemInfo.HardwareVersion,
		Led:             ledState,
		Power:           powerState,
	}

	return system
}
