package models

type ModulationData struct {
	Bandwidth int `json:"bandwidth"`
	Spreading int `json:"spreading"`
}
type HardwareData struct {
	Snr  int `json:"snr"`
	Rssi int `json:"rssi"`
}
type RadioData struct {
	Gps_time   int     		`json:"gps_time"`
	Hardware   HardwareData   	`json:"hardware"`
	DataRate   int        		`json:"datarate"`
	Modulation ModulationData 	`json:"modulation"`
	Delay      float64     		`json:"delay"`
	Freq       float64    		`json:"freq"`
	Size       int     		`json:"size"`
}


type Received struct {
	First AntaresDetail 	`json:"m2m:cin"`
}

type AntaresDetail struct {
	Rn  string  		`json:"rn"`
	Ty  int     		`json:"ty"`
	Ri  string  		`json:"ri"`
	Pi  string  		`json:"pi"`
	Ct  string  		`json:"ct"`
	Lt  string  		`json:"lt"`
	St  int     		`json:"st"`
	Cnf string  		`json:"cnf"`
	Cs  int     		`json:"cs"`
	Con string	`json:"con"`
}

type ConnectionDat struct {
	Type    string       	`json:"type"`
	Port    int          	`json:"port"`
	Data    DataReceived 	`json:"data"`
	Counter int          	`json:"counter"`
	DevEUI  string       	`json:"devEui"`
	Radio   RadioData		`json:"radio"`
}

type DataReceived struct {
	Volt float64 `json:"volt"`
	Cur  float64 `json:"cur"`
	Pwr  float64 `json:"pwr"`
}

