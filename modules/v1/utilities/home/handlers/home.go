package home

import (
	"github.com/gin-gonic/gin"
	"github.com/mhilmi999/Tugas-Akhir-Smart-Street-Light-LoRaWAN/modules/v1/utilities/home/repository"
	"github.com/mhilmi999/Tugas-Akhir-Smart-Street-Light-LoRaWAN/modules/v1/utilities/home/service"
	"gorm.io/gorm"
)

type HomeHandler interface{
	ReceivedData(c *gin.Context)
	SubscribeWebhook(c *gin.Context)
}

type homeHandler struct{
	homeService	service.Service
}

func NewHomeHandler(homeService service.Service) *homeHandler{
	return &homeHandler{homeService}
}

func Handler(db *gorm.DB) *homeHandler{
	Repository := repository.NewRepository(db)
	Service := service.NewService(Repository)
	return NewHomeHandler(Service)
}

