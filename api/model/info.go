package model

type Switch string

const (
	On  Switch = "on"
	Off        = "off"
)

type Info struct {
	Ip              string `json:"ip"`
	Port            int    `json:"port"`
	Name            string `json:"alias"`
	Icon            string `json:"icon_hash"`
	Model           string `json:"model"`
	MacAddress      string `json:"mac"`
	SoftwareVersion string `json:"softwareVersion"`
	HardwareVersion string `json:"hardwareVersion"`
	Led             Switch `json:"led"`
	Switch          Switch `json:"switch"`
}
