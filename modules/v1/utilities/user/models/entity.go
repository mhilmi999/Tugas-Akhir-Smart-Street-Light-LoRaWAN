package models

import "time"

type User struct {
	User_id      int
	Name         string
	Email        string
	Password     string
	Avatar       string
	Date_updated time.Time
	Date_created time.Time
}

type Modulation struct {
	Bandwidth int `json:"bandwidth"`
	Spreading int `json:"spreading"`
}

type Hardware struct {
	Snr  float64 `json:"snr"`
	Rssi float64 `json:"rssi"`
}

type Radio struct {
	Gps_time   string     `json:"gps_time"`
	Hardware   Hardware   `json:"hardware"`
	DataRate   int        `json:"datarate"`
	Modulation Modulation `json:"modulation"`
	Delay      string     `json:"delay"`
	Freq       float64    `json:"freq"`
	Size       string     `json:"size"`
}

type ConData struct {
	Type    string `json:"type"`
	Port    int    `json:"port"`
	Counter int    `json:"counter"`
	Data    string `json:"data"`
	Radio   Radio  `json:"radio"`
}
