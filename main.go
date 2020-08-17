package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net"

	"tp-link-hs110-api/crypto"
)

const (
	printDebug    = false
	StopCharacter = "\r\n\r\n"
	emeter        = "{\"emeter\":{\"get_realtime\":null}}"
	info          = "{\"system\":{\"get_sysinfo\":null}}"
	on            = "{\"system\":{\"set_relay_state\":{\"state\":1}}}}"
	off           = "{\"system\":{\"set_relay_state\":{\"state\":0}}}}"
)

var (
	encryptor = crypto.NewEncryptor(printDebug)
	decryptor = crypto.NewDecryptor(printDebug)
)

func main() {
	fmt.Println("Hello tp-link-hs110-api!")

	conn, err := net.Dial("tcp", "10.0.0.211:9999")

	if err != nil {
		log.Fatalln(err)
	}

	defer conn.Close()

	request := info

	conn.Write(encryptor.Encrypt(request))
	conn.Write([]byte(StopCharacter))

	log.Printf("Request: %s", request)

	all, _ := ioutil.ReadAll(conn)

	log.Printf("Response: %s", decryptor.Decrypt(all))
}
