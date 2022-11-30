package repository

import (
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
	BindSensorData(input models.ConnectionDat, DeviceId string) error
	GetChartData() ([]models.DeviceChartData, error)
	GetListDevice() ([]models.ListDevice, error)
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

func (n *repository) BindSensorData(input models.ConnectionDat, DeviceId string) error {
	err := n.db.Exec("INSERT INTO device_history (device_id, power, voltage, ampere, device_cons, history_date) VALUES (?,?,?,?,?,?)", DeviceId, input.Data.Pwr, input.Data.Volt, input.Data.Cur, 1, time.Now()).Error
	err = n.db.Exec("UPDATE device_monitoring SET power = ?, voltage = ?, ampere = ?, device_cons = ? WHERE device_id = ?", input.Data.Pwr, input.Data.Volt, input.Data.Cur, 1, DeviceId).Error
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
