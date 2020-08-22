package client

import (
	"bufio"
	"bytes"
	"io"
	"log"
	"net"
	"time"

	"tp-link-hs110-api/client/crypto"
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

	received, err := Read(conn)

	if err != nil {
		return "", err
	}

	if d.printDebug {
		log.Printf("Received: %s\n", received)
	}

	return received, nil
}

func Read(conn net.Conn) (string, error) {
	reader := bufio.NewReader(conn)
	var key = int32(0x2B)
	var buffer bytes.Buffer
	var counter = 0
	var countOpenBrackets = 1
	var countCloseBrackets = 0
	buffer.WriteString("{")

	for {
		ba, err := reader.ReadByte()
		if err != nil {
			if err == io.EOF {
				break
			}
			return "", err
		}

		decryptedValue := string(int32(ba) ^ key)
		key = int32(ba)

		// ignore first 4 bytes
		if counter > 4 {
			buffer.WriteString(decryptedValue)
		}

		// count opened brackets
		if decryptedValue == "{" {
			countOpenBrackets++
		}

		// count closed bracket
		if decryptedValue == "}" {
			countCloseBrackets++
		}

		// stop reading connection if opened and closed brackets are equal
		if countOpenBrackets == countCloseBrackets {
			break
		}

		counter++
	}
	return buffer.String(), nil
}
