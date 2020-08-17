package main

import (
	"encoding/binary"
	"fmt"
	"log"
	"net"
)

const (
	message       = "{\"emeter\":{\"get_realtime\":null}}"
	StopCharacter = "\r\n\r\n"
	on            = "{\"system\":{\"set_relay_state\":{\"state\":1}}}}"
	off           = "{\"system\":{\"set_relay_state\":{\"state\":0}}}}"
)

func main() {
	fmt.Println("Hello tp-link-hs110-api!")

	conn, err := net.Dial("tcp", "10.0.0.210:9999")

	if err != nil {
		log.Fatalln(err)
	}

	defer conn.Close()

	command := on

	data := encrypt(command)

	var text []byte
	text = append(text, 0)
	text = append(text, 0)
	text = append(text, 0)
	text = append(text, 0)
	binary.BigEndian.PutUint32(text, uint32(len(command)))

	for _, x := range data {
		valur := byte(x)
		//t := int8(x - 256)
		fmt.Printf("write %d to byte %b\n", x, valur)
		//fmt.Println(t)
		text = append(text, valur)
		//text = append(text, uint8(x-256))
	}

	conn.Write(text)
	conn.Write([]byte(StopCharacter))
	log.Printf("Send: %s", command)
}

func encrypt(message string) []int32 {
	var buffer []int32

	var key = int32(0xAB)

	for pos, char := range message {
		value := char ^ key
		key = value
		buffer = append(buffer, value)
		fmt.Printf("character %c, %d starts at byte position %d\n", char, value, pos)
	}

	return buffer
}
