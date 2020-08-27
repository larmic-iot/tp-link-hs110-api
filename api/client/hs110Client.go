package client

import (
	"encoding/json"
	"fmt"
	"log"
	"net"
	"strconv"
	"time"

	"tp-link-hs110-api/api/client/crypto"
	"tp-link-hs110-api/api/client/ioutil"
	"tp-link-hs110-api/api/client/model"
)

type TpLinkHS110Client struct {
	ip         string
	port       int
	timoutInMs int
	printDebug bool
}

const (
	// see https://github.com/softScheck/tplink-smartplug/blob/master/tplink-smarthome-commands.txt
	emeter        = "{\"emeter\":{\"get_realtime\":{}}}"
	dailyEMeter   = "{\"emeter\":{\"get_daystat\":{\"month\":%d,\"year\":%d}}}"
	info          = "{\"system\":{\"get_sysinfo\":{}}}"
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

func (d *TpLinkHS110Client) RequestInfo() (model.SystemInfo, error) {
	response, err := d.request(info)

	var wrapper model.SystemWrapper

	if err != nil {
		return model.SystemInfo{}, err
	}

	err = json.Unmarshal([]byte(response), &wrapper)

	if err != nil {
		return model.SystemInfo{}, err
	}

	return wrapper.System.SystemInfo, err
}

func (d *TpLinkHS110Client) RequestCurrentEnergyStatistics() (model.EMeterInfo, error) {
	response, err := d.request(emeter)

	if err != nil {
		return model.EMeterInfo{}, err
	}

	var wrapper model.EMeterWrapper
	err = json.Unmarshal([]byte(response), &wrapper)

	if err != nil {
		return model.EMeterInfo{}, err
	}

	return wrapper.EMeter.EMeterInfo, err
}

func (d *TpLinkHS110Client) RequestDailyEnergyStatistics(year int, month time.Month, day int) (model.DayStatEMeterInfo, error) {
	response, err := d.request(fmt.Sprintf(dailyEMeter, month, year))

	if err != nil {
		return model.DayStatEMeterInfo{}, err
	}

	var wrapper model.DayStatEMeterWrapper
	err = json.Unmarshal([]byte(response), &wrapper)

	if err != nil {
		return model.DayStatEMeterInfo{}, err
	}

	for _, value := range wrapper.DayStatEMeter.DayStatEMeterInfos.DayStatEMeterInfos {
		if value.Day == day {
			return value, nil
		}
	}

	return model.DayStatEMeterInfo{}, nil
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
	conn, err := dialer.Dial("tcp", d.ip+":"+strconv.Itoa(d.port))

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
