package view

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-contrib/sessions"
	"github.com/mhilmi999/Tugas-Akhir-Smart-Street-Light-LoRaWAN/modules/v1/utilities/home/repository"
	"github.com/mhilmi999/Tugas-Akhir-Smart-Street-Light-LoRaWAN/modules/v1/utilities/home/service"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type homeView struct {
	homeService service.Service
}

var hari = [...]string{
	"Minggu", "Senin", "Selasa", "Rabu", "Kamis", "Jumat", "Sabtu"}

var bulan = [...]string{
	"Jan", "Feb", "Mar", "Apr", "Mei", "Jun",
	"Jul", "Agu", "Sep", "Okt", "Nov", "Des",
}

func NewHomeView(homeService service.Service) *homeView {
	return &homeView{homeService}
}

func View(db *gorm.DB) *homeView {
	Repository := repository.NewRepository(db)
	Service := service.NewService(Repository)
	return NewHomeView(Service)
}

func Tanggal(t time.Time) string {
	return fmt.Sprintf("%s, %02d %s | %02d:%02d",
		hari[t.Weekday()], t.Day(), bulan[t.Month()-1][:3], t.Hour(), int(t.Minute()),
	)
}

func (h *homeView) Index(c *gin.Context) {
	session := sessions.Default(c)
	c.HTML(http.StatusOK, "home", gin.H{
		"title":				"Home SSL",
		"timeNow": 				Tanggal(time.Now()),
		"UserID":				session.Get("userID"),
		"UserName":				session.Get("userName"),
	})
}

func (h *homeView) Login(c *gin.Context) {
	c.HTML(http.StatusOK, "login", gin.H{
		"title": "Login SSL",
	})
}

func (h *homeView) Register(c *gin.Context) {
	c.HTML(http.StatusOK, "register", gin.H{
		"title": "Register SSL",
	})
}
