package client

import (
	"io/ioutil"
	"log"
	"net"

	"tp-link-hs110-api/client/crypto"
)

type TpLinkHS110Client struct {
	ip         string
	port       int
	printDebug bool
}

const (
	emeter        = "{\"emeter\":{\"get_realtime\":null}}"
	info          = "{\"system\":{\"get_sysinfo\":null}}"
	on            = "{\"system\":{\"set_relay_state\":{\"state\":1}}}}"
	off           = "{\"system\":{\"set_relay_state\":{\"state\":0}}}}"
	StopCharacter = "\r\n\r\n"
)

func NewTpLinkHS110Client(ip string, printDebug bool) *TpLinkHS110Client {
	return &TpLinkHS110Client{
		ip:         ip,
		port:       9999,
		printDebug: printDebug,
	}
}

func (d *TpLinkHS110Client) RequestInfo() string {
	return d.request(info)
}

func (d *TpLinkHS110Client) RequestCurrentEnergyStatistics() string {
	return d.request(emeter)
}

func (d *TpLinkHS110Client) RequestSwitchOn() string {
	return d.request(on)
}

func (d *TpLinkHS110Client) RequestSwitchOff() string {
	return d.request(off)
}

func (d *TpLinkHS110Client) request(message string) string {
	encryptor := crypto.NewEncryptor(d.printDebug)
	decryptor := crypto.NewDecryptor(d.printDebug)
	conn, err := net.Dial("tcp", d.ip+":9999")

	if err != nil {
		log.Fatalln(err)
	}

	defer conn.Close()

	if d.printDebug {
		log.Printf("Sent:     %s\n", message)
	}

	_, _ = conn.Write(encryptor.Encrypt(message))
	_, _ = conn.Write([]byte(StopCharacter))

	all, _ := ioutil.ReadAll(conn)

	received := decryptor.Decrypt(all)

	if d.printDebug {
		log.Printf("Received: %s\n", received)
	}

	return received
}
