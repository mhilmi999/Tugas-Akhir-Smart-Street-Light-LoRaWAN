package models

import "time"

type DeviceChartData struct {
	Voltage float64 
	Ampere  float64 
	Power  float64 
	History_date time.Time
}
