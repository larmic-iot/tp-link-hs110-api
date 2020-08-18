package main

import (
	"log"

	"tp-link-hs110-api/client"
)

const (
	printDebug = false
)

func main() {
	log.Println("Hello tp-link-hs110-api!")

	fridgeClient := client.NewTpLinkHS110Client("10.0.0.210", printDebug)

	log.Printf("Request info: %s", fridgeClient.RequestInfo())
	log.Printf("Request energy: %s", fridgeClient.RequestCurrentEnergyStatistics())
}
