package model

type EMeterWrapper struct {
	EMeter EMeter `json:"emeter"`
}

type EMeter struct {
	EMeterInfo EMeterInfo `json:"get_realtime"`
}

type EMeterInfo struct {
	Voltage         int64 `json:"voltage_mv"`
	CurrentMilliAmp int64 `json:"current_ma"`
	Power           int64 `json:"power_mw"`
	TotalWattHours  int64 `json:"total_wh"`
	ErrorCode       int64 `json:"err_code"`
}
