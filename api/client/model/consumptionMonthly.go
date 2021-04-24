package model

type MonthStatEMeterWrapper struct {
	MonthStatEMeter MonthStatEMeter `json:"emeter"`
}

type MonthStatEMeter struct {
	MonthStatEMeterInfos MonthStatEMeterInfos `json:"get_monthstat"`
}

type MonthStatEMeterInfos struct {
	MonthStatEMeterInfo []MonthStatEMeterInfo `json:"month_list"`
}

type MonthStatEMeterInfo struct {
	Year            int   `json:"year"`
	Month           int   `json:"month"`
	EnergyWattHours int64 `json:"energy_wh"`
}
