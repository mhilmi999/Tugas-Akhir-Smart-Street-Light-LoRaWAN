package service

import (
	// "encoding/json"
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
