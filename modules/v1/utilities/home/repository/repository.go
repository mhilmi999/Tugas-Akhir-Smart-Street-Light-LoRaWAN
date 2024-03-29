package repository

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/mhilmi999/Tugas-Akhir-Smart-Street-Light-LoRaWAN/modules/v1/utilities/home/models"
	"gorm.io/gorm"
)

type Repository interface {
	GetLatestCon(token string) (models.Received, error)
	CheckEnergyExist(antaresDeviceID string) (float64, error)
	BindSensorData(input models.ConnectionDat, DeviceId string, energyExist float64) error
	GetChartData() ([]models.DeviceChartData, error)
	GetListDevice() ([]models.ListDevice, error)
	ControlLight(power string, deviceid string, token string) (int, error)
	GetEstimatedPrice(deviceId string) (models.EstPrice, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (n *repository) GetLatestCon(token string) (models.Received, error) {
	data := models.Received{}

	client := http.Client{}
	req, err := http.NewRequest("GET", "https://platform.antares.id:8443/~/antares-cse/antares-id/SmartStreetLight/tugasAkhir/la", nil)
	req.Header.Set("X-M2M-Origin", token)
	req.Header.Set("Content-Type", "application/json;ty=4")
	req.Header.Set("Accept", "application/json")
	if err != nil {
		fmt.Println(err)
		return data, err
	}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return data, err
	}
	defer resp.Body.Close()
	isiBody, err := ioutil.ReadAll(resp.Body)
	json.Unmarshal(isiBody, &data)
	fmt.Println(data.First.Con)
	return data, err
}

func (n *repository) CheckEnergyExist(antaresDeviceID string) (float64, error) {
	var energyExist float64
	err := n.db.Raw("SELECT energy_total FROM device_monitoring WHERE device_id = ?", antaresDeviceID).Scan(&energyExist).Error
	return energyExist, err
}

func (n *repository) BindSensorData(input models.ConnectionDat, DeviceId string, energyExist float64) error {
	err := n.db.Exec("INSERT INTO device_history (device_id, power, voltage, ampere, device_cons, history_date, energy) VALUES (?,?,?,?,?,?,?)", DeviceId, input.Data.P, input.Data.V, input.Data.C, 1, time.Now(), input.Data.E).Error
	err = n.db.Exec("UPDATE device_monitoring SET power = ?, voltage = ?, ampere = ?, device_cons = ?, energy_total = ?, last_updated = ? WHERE device_id = ?", input.Data.P, input.Data.V, input.Data.C, 1, energyExist+input.Data.E, time.Now(), DeviceId).Error
	return err
}
func (n *repository) GetChartData() ([]models.DeviceChartData, error) {
	var chartData []models.DeviceChartData
	err := n.db.Raw("SELECT power, voltage, ampere, history_date FROM device_history").Scan(&chartData).Error
	return chartData, err
}

func (n *repository) GetListDevice() ([]models.ListDevice, error) {
	var listDevice []models.ListDevice
	err := n.db.Raw("SELECT d.device_id, m.voltage, m.ampere, m.power, m.device_cons FROM device d INNER JOIN device_monitoring m ON d.device_id = m.device_id;").Scan(&listDevice).Error
	return listDevice, err
}

func (n *repository) ControlLight(power string, deviceid string, token string) (int, error) {
	data := "\r\n{\r\n  \"m2m:cin\": {\r\n    \"con\": \r\n      \"{\r\n      \t \\\"type\\\":\\\"downlink\\\",\r\n      \\\"data\\\":\\\"" + power + "\\\"\r\n      }\"\r\n    }\r\n}"

	client := http.Client{}
	req, err := http.NewRequest("POST", "https://platform.antares.id:8443/~/antares-cse/antares-id/SmartStreetLight/tugasAkhir", bytes.NewBuffer([]byte(data)))
	// req, err := http.NewRequest("POST", "https://platform.antares.id:8443/~/antares-cse/antares-id/AntaresDevKitTest/HilmiUplinkTest", bytes.NewBuffer([]byte(data)))
	req.Header.Set("X-M2M-Origin", token)
	req.Header.Set("Content-Type", "application/json;ty=4")
	req.Header.Set("Accept", "application/json")
	if err != nil {
		fmt.Println(err)
		return 0, err
	}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return 0, err
	}
	defer resp.Body.Close()
	_, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return 0, err
	}
	n.db.Exec("UPDATE device_monitoring SET device_cons = ? WHERE device_id = ?", power, deviceid)
	return 1, err
}

func (n *repository) GetEstimatedPrice(deviceId string) (models.EstPrice, error) {
	var estPrice models.EstPrice
	err := n.db.Raw("SELECT energy_total, last_updated FROM device_monitoring WHERE device_id = ?", deviceId).Scan(&estPrice).Error
	estPrice.Price_Est = estPrice.Energy_total * 1699.53
	fmt.Println(estPrice)
	return estPrice, err
}
