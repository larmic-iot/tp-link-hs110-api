package model

// TODO map plug is on or off
type Info struct {
	Identifier      string `json:"Identifier"`
	Name            string `json:"alias"`
	Icon            string `json:"icon_hash"`
	Model           string `json:"model"`
	MacAddress      string `json:"mac"`
	SoftwareVersion string `json:"softwareVersion"`
	HardwareVersion string `json:"hardwareVersion"`
	Led             bool   `json:"led"`
}
