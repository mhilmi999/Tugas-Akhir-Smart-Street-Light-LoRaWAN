package user

import (
	"github.com/mhilmi999/Tugas-Akhir-Smart-Street-Light-LoRaWAN/modules/v1/utilities/user/repository"
	"github.com/mhilmi999/Tugas-Akhir-Smart-Street-Light-LoRaWAN/modules/v1/utilities/user/service"
	"gorm.io/gorm"
)

func Handler(db *gorm.DB) *userHandler {
	userRepository := repository.NewRepository(db)
	userService := service.NewService(userRepository)
	userHandler := NewUserHandler(userService)
	return userHandler
}