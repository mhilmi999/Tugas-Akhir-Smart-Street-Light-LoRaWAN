package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/mhilmi999/Tugas-Akhir-Smart-Street-Light-LoRaWAN/app/config"
	homeViewV1 "github.com/mhilmi999/Tugas-Akhir-Smart-Street-Light-LoRaWAN/modules/v1/utilities/home/view"
	"gorm.io/gorm"
)

func ParseTmpl(router *gin.Engine) *gin.Engine {
	router.Static("/assets", "./public/assets")
	router.Static("/images", "./public/assets/images")
	router.Static("/css", "./public/assets/css")
	router.Static("/js", "./public/assets/js")
	router.Static("/vendor", "./public/assets/vendor")
	return router
}

func Init(db *gorm.DB, conf config.Conf, router *gin.Engine) *gin.Engine {
	homeViewV1 := homeViewV1.View(db)
	// Routing to website service
	home := router.Group("/")
	home.GET("", homeViewV1.Index)
	home.GET("/home", homeViewV1.Index)
	router = ParseTmpl(router)

	return router
}
