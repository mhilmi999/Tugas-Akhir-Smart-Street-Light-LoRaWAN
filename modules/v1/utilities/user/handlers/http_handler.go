package user

import "github.com/mhilmi999/Tugas-Akhir-Smart-Street-Light-LoRaWAN/modules/v1/utilities/home/service"

type UserHandler interface{
}

type userHandler struct{
	userService service.Service
}

func NewUserHandler (userService service.Service) *userHandler{
	return &userHandler{userService}
}