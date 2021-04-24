package api

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strconv"
	"tp-link-hs110-api/api/client"
	http3 "tp-link-hs110-api/api/http"
	parser2 "tp-link-hs110-api/api/parser"
)

func ConsumptionHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	ipParameter := vars["ip"]

	ip, err := parser2.ParseIp(ipParameter)

	if err != nil {
		http3.NewErrorEncoder(w).Encode(http.StatusBadRequest, ipParameter+" not valid")
		return
	}

	year, _ := strconv.Atoi(vars["year"])
	month := vars["month"]
	day := vars["day"]

	// TODO implement me
	log.Printf("ip: %s", ip)
	log.Printf("year: %d", year)
	log.Printf("month: %s", month)
	log.Printf("day: %s", day)

	socketClient := client.NewTpLinkHS110Client(ipParameter, timeoutInMs, printDebug)

	atoi2, _ := strconv.Atoi(month)
	log.Println(socketClient.RequestMonthlyEnergyStatistics(year, atoi2))

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
}
