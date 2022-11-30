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
	
	fmt.Println("Ini hasil data sensornya \n",data.Data)
	fmt.Println("Ini hasil data Con fullnya \n",data)
	return data, nil
}

func (n *service) GetChartData() ([]models.DeviceChartData, error) {
	chartData, err := n.repository.GetChartData()
	return chartData, err
}