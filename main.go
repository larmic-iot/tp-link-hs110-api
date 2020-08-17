package main

import (
	"bytes"
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
)

func main() {
	fmt.Println("Hello tp-link-hs110-api!")

	conn, err := net.Dial("tcp", "10.0.0.210:9999")

	if err != nil {
		log.Fatalln(err)
	}

	defer conn.Close()

	command := info

	conn.Write(encryptor.Encrypt(command))
	conn.Write([]byte(StopCharacter))

	log.Printf("Send: %s", command)

	all, _ := ioutil.ReadAll(conn)
	log.Println(all)

	var key = int32(0x2B)

	var buffer bytes.Buffer
	buffer.WriteString("{")

	for pos, char := range all {
		value := int32(char) ^ key
		key = int32(char)
		fmt.Printf("read character %d, %d, %s\n", int32(char), value, string(value))
		if pos > 4 {
			buffer.WriteString(string(value))
		}
	}

	fmt.Println(buffer.String())
}
