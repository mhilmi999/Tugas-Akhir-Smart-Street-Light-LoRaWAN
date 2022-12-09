package models

type ModulationData struct {
	Bandwidth int `json:"bandwidth"`
	Spreading int `json:"spreading"`
}
type HardwareData struct {
	Snr  float64 `json:"snr"`
	Rssi int     `json:"rssi"`
}
type RadioData struct {
	Gps_time   int            `json:"gps_time"`
	Hardware   HardwareData   `json:"hardware"`
	DataRate   int            `json:"datarate"`
	Modulation ModulationData `json:"modulation"`
	Delay      float64        `json:"delay"`
	Freq       float64        `json:"freq"`
	Size       int            `json:"size"`
}

type Received struct {
	First AntaresDetail `json:"m2m:cin"`
}

type AntaresDetail struct {
	Rn  string `json:"rn"`
	Ty  int    `json:"ty"`
	Ri  string `json:"ri"`
	Pi  string `json:"pi"`
	Ct  string `json:"ct"`
	Lt  string `json:"lt"`
	St  int    `json:"st"`
	Cnf string `json:"cnf"`
	Cs  int    `json:"cs"`
	Con string `json:"con"`
}

type ConnectionDat struct {
	Type    string       `json:"type"`
	Port    int          `json:"port"`
	Data    DataReceived `json:"data"`
	Counter int          `json:"counter"`
	DevEUI  string       `json:"devEui"`
	Radio   RadioData    `json:"radio"`
}

type DataReceived struct {
	V float64 `json:"v"`
	C  float64 `json:"c"`
	P  float64 `json:"p"`
	E  float64 `json:"e"`
}

type ObjectAntares5 struct {
	Rn  string `json:"rn"`
	Ty  int    `json:"ty"`
	Ri  string `json:"ri"`
	Pi  string `json:"pi"`
	Ct  string `json:"ct"`
	Lt  string `json:"lt"`
	St  int    `json:"st"`
	Cnf string `json:"cnf"`
	Cs  int    `json:"cs"`
	Con string `json:"con"`
}

type ObjectAntares4 struct {
	M2m_cin ObjectAntares5 `json:"m2m:cin"`
}

type ObjectAntares3 struct {
	M2m_rep ObjectAntares4 `json:"m2m:rep"`
	M2m_rss int            `json:"m2m:rss"`
}

type ObjectAntares2 struct {
	M2m_nev ObjectAntares3 `json:"m2m:nev"`
	M2m_sud bool           `json:"m2m:sud"`
	M2m_sur string         `json:"m2m:sur"`
}

type ObjectAntares1 struct {
	First ObjectAntares2 `json:"m2m:sgn"`
}
