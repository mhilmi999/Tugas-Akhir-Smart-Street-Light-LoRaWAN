package main

import (
	"log"

	"github.com/mhilmi999/Tugas-Akhir-Smart-Street-Light-LoRaWAN/app/config"
	"github.com/mhilmi999/Tugas-Akhir-Smart-Street-Light-LoRaWAN/app/database"
	routesV1 "github.com/mhilmi999/Tugas-Akhir-Smart-Street-Light-LoRaWAN/modules/v1/routes"
	"github.com/mhilmi999/Tugas-Akhir-Smart-Street-Light-LoRaWAN/pkg/html"

	// "github.com/mhilmi999/Tugas-Akhir-Smart-Street-Light-LoRaWAN/pkg/http-error"

	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"

	// "github.com/gin-contrib/sessions"
	// "github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func setup() (*gorm.DB, config.Conf, *gin.Engine){
	conf, err := config.Init()
	gin.SetMode(conf.App.Mode)
	if err != nil {
		log.Fatal(err)
	}
	db := database.Init(conf)
	router := gin.Default()
	router.Use(cors.Default())

	cookieStore := cookie.NewStore([]byte(conf.App.Secret_key))
	router.Use(sessions.Sessions("homesmartstreetlight", cookieStore))
	router.HTMLRender = html.Render("./public/templates")

	//Error Handling for 404 Not Found Page and Method Not Allowed
	// router.NoRoute(error.PageNotFound())
	// router.NoMethod(error.NoMethod())
	return db, conf, router
}

func main() {
	conf, err := config.Init()
	if err != nil {
		log.Fatal(err)
	}

	router := routesV1.Init(setup())

	router.Run(":" + conf.App.Port)
}