package models

import "time"

type DeviceChartData struct {
	Voltage float64 
	Ampere  float64 
	Power  float64 
	History_date time.Time
}

type ListDevice struct{
	Device_id string
	Voltage float64
	Ampere float64
	Power float64
	Device_cons int
}