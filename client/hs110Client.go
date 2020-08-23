package client

import (
	"encoding/json"
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

type SystemWrapper struct {
	System System `json:"system"`
}

type System struct {
	SystemInfo SystemInfo `json:"get_sysinfo"`
}

type SystemInfo struct {
	SoftwareVersion string  `json:"sw_ver"`
	HardwareVersion string  `json:"hw_ver"`
	Model           string  `json:"model"`
	DeviceId        string  `json:"deviceId"`
	OemId           string  `json:"oemId"`
	HardwareId      string  `json:"hwId"`
	RSSI            int64   `json:"rssi"`
	Longitude       float64 `json:"longitude_i"`
	Latitude        float64 `json:"latitude_i"`
	Alias           string  `json:"alias"`
	Status          string  `json:"status"`
	MicType         string  `json:"mic_type"`
	Feature         string  `json:"feature"`
	MacAddress      string  `json:"mac"`
	Updating        int     `json:"updating"`
	LedOff          int     `json:"led_off"`
	RelayState      int     `json:"relay_state"`
	OnTime          int64   `json:"on_time"`
	ActiveMode      string  `json:"active_mode"`
	Icon            string  `json:"icon_hash"`
	DevName         string  `json:"dev_name"`
	ErrorCode       int     `json:"err_code"`
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

func (d *TpLinkHS110Client) RequestInfo() (SystemInfo, error) {
	response, err := d.request(info)

	var systemWrapper SystemWrapper

	if err != nil {
		return SystemInfo{}, err
	}

	err = json.Unmarshal([]byte(response), &systemWrapper)

	if err != nil {
		return SystemInfo{}, err
	}

	return systemWrapper.System.SystemInfo, err
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
