package service

import (
	"encoding/json"
	"fmt"

	"github.com/mhilmi999/Tugas-Akhir-Smart-Street-Light-LoRaWAN/modules/v1/utilities/home/models"
	"github.com/mhilmi999/Tugas-Akhir-Smart-Street-Light-LoRaWAN/modules/v1/utilities/home/repository"
)

type Service interface{	
	GetLatestCon(token string)(models.Received, error)
	GetDatafromCon(input string)(models.ConnectionDat, error)
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

func(n *service) GetDatafromCon(input string)(models.ConnectionDat, error){
	var data models.ConnectionDat
	err := json.Unmarshal([]byte(input), &data)
	if err != nil{
		fmt.Println(err)
		return data, err
	}
	fmt.Println("Ini di service",data.Data)
	return data, nil
}