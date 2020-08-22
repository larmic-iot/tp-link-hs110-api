package client

import (
	"log"
	"net"
	"time"

	"tp-link-hs110-api/client/crypto"
	"tp-link-hs110-api/client/ioutil"
)

type TpLinkHS110Client struct {
	ip         string
	port       int
	timoutInMs int
	printDebug bool
}

const (
	// see https://github.com/softScheck/tplink-smartplug/blob/master/tplink-smarthome-commands.txt
	emeter        = "{\"emeter\":{\"get_realtime\":null}}"
	info          = "{\"system\":{\"get_sysinfo\":null}}"
	on            = "{\"system\":{\"set_relay_state\":{\"state\":1}}}}"
	off           = "{\"system\":{\"set_relay_state\":{\"state\":0}}}}"
	StopCharacter = "\r\n\r\n"
)

func NewTpLinkHS110Client(ip string, timoutInMs int, printDebug bool) *TpLinkHS110Client {
	return &TpLinkHS110Client{
		ip:         ip,
		port:       9999,
		timoutInMs: timoutInMs,
		printDebug: printDebug,
	}
}

func (d *TpLinkHS110Client) RequestInfo() (string, error) {
	return d.request(info)
}

func (d *TpLinkHS110Client) RequestCurrentEnergyStatistics() (string, error) {
	return d.request(emeter)
}

func (d *TpLinkHS110Client) RequestSwitchOn() (string, error) {
	return d.request(on)
}

func (d *TpLinkHS110Client) RequestSwitchOff() (string, error) {
	return d.request(off)
}

func (d *TpLinkHS110Client) request(message string) (string, error) {
	encryptor := crypto.NewEncryptor(d.printDebug)
	dialer := net.Dialer{Timeout: time.Duration(d.timoutInMs) * time.Millisecond}
	conn, err := dialer.Dial("tcp", d.ip+":9999")

	if err != nil {
		return "", err
	}

	defer conn.Close()

	if d.printDebug {
		log.Printf("Sent:     %s\n", message)
	}

	_, err = conn.Write(encryptor.Encrypt(message))
	_, err = conn.Write([]byte(StopCharacter))

	if err != nil {
		return "", err
	}

	received, err := ioutil.ReadJson(conn)

	if err != nil {
		return "", err
	}

	if d.printDebug {
		log.Printf("Received: %s\n", received)
	}

	return received, nil
}
