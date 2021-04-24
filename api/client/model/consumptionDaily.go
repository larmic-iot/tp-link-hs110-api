package model

type DayStatEMeterWrapper struct {
	DayStatEMeter DayStatEMeter `json:"emeter"`
}

type DayStatEMeter struct {
	DayStatEMeterInfos DayStatEMeterInfos `json:"get_daystat"`
}

type DayStatEMeterInfos struct {
	DayStatEMeterInfos []DayStatEMeterInfo `json:"day_list"`
}

type DayStatEMeterInfo struct {
	Year            int   `json:"year"`
	Month           int   `json:"month"`
	Day             int   `json:"day"`
	EnergyWattHours int64 `json:"energy_wh"`
}
