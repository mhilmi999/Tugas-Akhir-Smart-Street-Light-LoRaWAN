package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/mhilmi999/Tugas-Akhir-Smart-Street-Light-LoRaWAN/app/config"

	"github.com/mhilmi999/Tugas-Akhir-Smart-Street-Light-LoRaWAN/app/middlewares"
	//mqttService "github.com/mhilmi999/Tugas-Akhir-Smart-Street-Light-LoRaWAN/embeddedSys/protocols/mqtt"
	homeHandlerV1 "github.com/mhilmi999/Tugas-Akhir-Smart-Street-Light-LoRaWAN/modules/v1/utilities/home/handlers"
	homeRepositoryV1 "github.com/mhilmi999/Tugas-Akhir-Smart-Street-Light-LoRaWAN/modules/v1/utilities/home/repository"
	homeServiceV1 "github.com/mhilmi999/Tugas-Akhir-Smart-Street-Light-LoRaWAN/modules/v1/utilities/home/service"
	homeViewV1 "github.com/mhilmi999/Tugas-Akhir-Smart-Street-Light-LoRaWAN/modules/v1/utilities/home/view"
	userHandlerV1 "github.com/mhilmi999/Tugas-Akhir-Smart-Street-Light-LoRaWAN/modules/v1/utilities/user/handlers"
	userRepositoryV1 "github.com/mhilmi999/Tugas-Akhir-Smart-Street-Light-LoRaWAN/modules/v1/utilities/user/repository"
	userServiceV1 "github.com/mhilmi999/Tugas-Akhir-Smart-Street-Light-LoRaWAN/modules/v1/utilities/user/service"
	pkgHtml "github.com/mhilmi999/Tugas-Akhir-Smart-Street-Light-LoRaWAN/pkg/html"
	"gorm.io/gorm"
)

func ParseTmpl(router *gin.Engine) *gin.Engine {
	router.Static("/assets", "./public/assets")
	router.Static("/images", "./public/assets/images")
	router.Static("/css", "./public/assets/css")
	router.Static("/js", "./public/assets/js")
	router.Static("/vendor", "./public/assets/vendor")
	router.HTMLRender = pkgHtml.ManualRender("./public/templates/")
	return router
}

func Init(db *gorm.DB, conf config.Conf, router *gin.Engine) *gin.Engine {
	homeRepositoryV1 := homeRepositoryV1.NewRepository(db)
	homeServiceV1 := homeServiceV1.NewService(homeRepositoryV1)
	homeHandlerV1 := homeHandlerV1.NewHomeHandler(homeServiceV1)
	homeViewV1 := homeViewV1.View(db)
	userRepositoryV1 := userRepositoryV1.NewRepository(db)
	userServiceV1 := userServiceV1.NewService(userRepositoryV1)
	userHandlerV1 := userHandlerV1.NewUserHandler(userServiceV1)

	// Mqtt using go routine (mini thread)
	//go mqttService.MqttRun()

	// Routing to website service
	home := router.Group("/")
	home.GET("", middlewares.IsLogin(), homeViewV1.Index)
	home.GET("/home", middlewares.IsLogin(), homeViewV1.Index)
	// home.GET("", homeViewV1.Index)
	// home.GET("/home", homeViewV1.Index)
	home.GET("/login", homeViewV1.Login)
	home.GET("/register", homeViewV1.Register)
	home.GET("/logout", userHandlerV1.Logout)
	home.GET("/list-device", middlewares.IsLogin(), homeViewV1.ListDevice)
	// home.GET("/list-device", homeViewV1.ListDevice)
	home.GET("/antares-data", homeHandlerV1.ReceivedData)
	home.POST("/login", userHandlerV1.Login)
	home.POST("register", userHandlerV1.Register)
	home.POST("/webhook", homeHandlerV1.SubscribeWebhook)
	home.GET("/manual-control/:power/:deviceid", middlewares.IsLogin(), homeHandlerV1.ControlLight)
	// home.GET("/manual-control/:power/:deviceid", homeHandlerV1.ControlLight)
	router = ParseTmpl(router)

	return router
}
