package client

import (
	"io/ioutil"
	"log"
	"net"

	"tp-link-hs110-api/crypto"
)

type TpLinkHS110Client struct {
	ip         string
	port       int
	printDebug bool
}

const (
	StopCharacter = "\r\n\r\n"
)

func NewTpLinkHS110Client(ip string, printDebug bool) *TpLinkHS110Client {
	return &TpLinkHS110Client{
		ip:         ip,
		port:       9999,
		printDebug: printDebug,
	}
}

func (d *TpLinkHS110Client) Request(message string) string {
	encryptor := crypto.NewEncryptor(d.printDebug)
	decryptor := crypto.NewDecryptor(d.printDebug)
	conn, err := net.Dial("tcp", d.ip+":9999")

	if err != nil {
		log.Fatalln(err)
	}

	defer conn.Close()

	request := message

	conn.Write(encryptor.Encrypt(request))
	conn.Write([]byte(StopCharacter))

	all, _ := ioutil.ReadAll(conn)

	return decryptor.Decrypt(all)
}
