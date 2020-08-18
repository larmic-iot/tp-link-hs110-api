package main

import (
	"log"

	"tp-link-hs110-api/client"
)

const (
	printDebug = false
	emeter     = "{\"emeter\":{\"get_realtime\":null}}"
	info       = "{\"system\":{\"get_sysinfo\":null}}"
	on         = "{\"system\":{\"set_relay_state\":{\"state\":1}}}}"
	off        = "{\"system\":{\"set_relay_state\":{\"state\":0}}}}"
)

func main() {
	log.Println("Hello tp-link-hs110-api!")

	fridgeClient := client.NewTpLinkHS110Client("10.0.0.210", printDebug)

	request := info

	response := fridgeClient.Request(request)

	log.Printf("Request: %s", request)
	log.Printf("Response: %s", response)
}
