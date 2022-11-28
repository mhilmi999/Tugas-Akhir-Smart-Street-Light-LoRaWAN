package service

import (
	// "encoding/json"
	"errors"
	"fmt"
	"time"

	"github.com/mhilmi999/Tugas-Akhir-Smart-Street-Light-LoRaWAN/modules/v1/utilities/user/models"
	"github.com/mhilmi999/Tugas-Akhir-Smart-Street-Light-LoRaWAN/modules/v1/utilities/user/repository"
	"golang.org/x/crypto/bcrypt"
)

type Service interface {
	Register(input models.RegisterInput) (models.User, error)
	Login(input models.LoginInput) (models.User, error)
	// GetId(id string) int
	// GetData(input string) string
	// DataCheck(id int, getData string) int
	GetLatestData(token string)(models.Received, error)
}

type service struct {
	repository repository.Repository
}

func NewService(repository repository.Repository) *service {
	return &service{repository}
}

func (n *service) Register(input models.RegisterInput) (models.User, error) {
	user := models.User{}
	user.Name = input.Nama
	user.Email = input.Email
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.MinCost)
	if err != nil {
		return user, err
	}
	user.Password = string(passwordHash)
	user.Avatar = "Images/avatar.png"
	user.Date_created = time.Now()
	user.Date_updated = time.Now()
	newUser, err := n.repository.Save(user)
	if err != nil {
		return newUser, err
	}
	return newUser, nil
}

func (n *service) Login(input models.LoginInput) (models.User, error) {
	email := input.Email
	password := input.Password
	user, _ := n.repository.FindByEmail(email)

	if user.User_id == 0 {
		return user, errors.New("Email yang Anda masukan salah")
	}

	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		fmt.Println(err)
		return user, errors.New("Password yang Anda masukan salah/tidak terdaftar")
	}
	return user, nil
}

func (n *service) GetLatestData(token string)(models.Received, error){
	getLatestData, err := n.repository.GetLatestData(token)
	return getLatestData, err
}

// func (n *service) GetId(id string) int {
// 	idData, err := n.repository.GetId(id)
// 	if err != nil || idData == 0 {
// 		fmt.Println(err)
// 		return idData
// 	}
// 	return idData
// }

// func (n *service) GetData(input string) string {
// 	fmt.Println(input)
// 	var Recdata models.ConData
// 	var err = json.Unmarshal([]byte(input), &Recdata)
// 	if err != nil {
// 		fmt.Println(err)
// 		return Recdata.Data
// 	}
// 	return Recdata.Data
// }

// func (n *service) DataCheck(id int, getData string) int {
// 	var re int
// 	fmt.Println(getData)
// 	re = n.SensorData(id, getData)

// 	return re
// }

// func (n *service) SensorData(id int, getData string) int {

// 	return 1
// }
