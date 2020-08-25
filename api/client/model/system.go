package model

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
