package service

import (
	"encoding/json"
	"fmt"

	"github.com/mhilmi999/Tugas-Akhir-Smart-Street-Light-LoRaWAN/modules/v1/utilities/home/models"
	"github.com/mhilmi999/Tugas-Akhir-Smart-Street-Light-LoRaWAN/modules/v1/utilities/home/repository"
)

type Service interface{	
	GetLatestCon(token string)(models.Received, error)
	GetDatafromCon(input string, DeviceId string)(models.ConnectionDat, error)
	GetChartData() ([]models.DeviceChartData, error) 
	GetListDevice()([]models.ListDevice, error)
	GetDatafromWebhook(sensorData string, antaresDeviceID string)(models.ConnectionDat, error)
	ControlLight(power string, token string)int
}

type service struct{
	repository repository.Repository
}

func NewService(repository repository.Repository) *service{
	return &service{repository}
}

func (n *service) GetLatestCon(token string)(models.Received, error){
	getLatestData, err := n.repository.GetLatestCon(token)
	return getLatestData, err
}

func(n *service) GetDatafromCon(input string, DeviceId string)(models.ConnectionDat, error){
	var data models.ConnectionDat
	err := json.Unmarshal([]byte(input), &data)
	if err != nil{
		fmt.Println(err)
		return data, err
	}
	err = n.repository.BindSensorData(data, DeviceId)
	
	return data, nil
}

func (n *service) GetChartData() ([]models.DeviceChartData, error) {
	chartData, err := n.repository.GetChartData()
	return chartData, err
}

func (n *service) GetListDevice()([]models.ListDevice, error){
	listDevice, err := n.repository.GetListDevice()
	return listDevice, err
}

func (n *service) GetDatafromWebhook(sensorData string, antaresDeviceID string)(models.ConnectionDat, error){
	var data models.ConnectionDat
	err := json.Unmarshal([]byte(sensorData), &data)
	if err != nil{
		fmt.Println(err)
		return data,err
	}
	fmt.Println("Ini hasil data sensor arusnya \n",data.Data)
	err = n.repository.BindSensorData(data, antaresDeviceID)
	return data,err
}

func (n *service) ControlLight(power string, token string)int{
	rtn, err := n.repository.ControlLight(power, token)
	if err != nil{
		fmt.Println(err)
		return 0
	}
	return rtn
}