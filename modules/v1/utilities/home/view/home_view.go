package view

import (
	"github.com/mhilmi999/Tugas-Akhir-Smart-Street-Light-LoRaWAN/modules/v1/utilities/home/repository"
	"github.com/mhilmi999/Tugas-Akhir-Smart-Street-Light-LoRaWAN/modules/v1/utilities/home/service"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type homeView struct {
	homeService service.Service
}

func NewHomeView(homeService service.Service) *homeView {
	return &homeView{homeService}
}

func View(db *gorm.DB) *homeView{
	Repository := repository.NewRepository(db)
	Service := service.NewService(Repository)
	return NewHomeView(Service)
}

func (h *homeView) Index(c *gin.Context){
	c.HTML(http.StatusOK, "home_index.html", nil)
}