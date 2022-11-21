package user

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/mhilmi999/Tugas-Akhir-Smart-Street-Light-LoRaWAN/modules/v1/utilities/user/models"
	"github.com/mhilmi999/Tugas-Akhir-Smart-Street-Light-LoRaWAN/modules/v1/utilities/user/service"
)

type UserHandler interface {
	Register(c *gin.Context)
	Login(c *gin.Context)
}

type userHandler struct {
	userService service.Service
}

func NewUserHandler(userService service.Service) *userHandler {
	return &userHandler{userService}
}

func (n *userHandler) Register(c *gin.Context) {
	var input models.RegisterInput
	fmt.Println(input)
	err := c.ShouldBind(&input)
	_, err = n.userService.Register(input)
	if err != nil {
		c.Redirect(http.StatusFound, "/register")
		fmt.Println(err)
		return
	}
	fmt.Println("Sukses Register")
	c.Redirect(http.StatusFound, "/login")
}

func (n *userHandler) Login(c *gin.Context) {
	session := sessions.Default(c)
	var input models.LoginInput
	err := c.ShouldBind(&input)
	if err != nil {
		log.Println(err)
		session.AddFlash("Format email salah")
		c.HTML(http.StatusOK, "login", gin.H{
			"title":    "Login SSL",
			"flashses": session.Flashes(),
		})
		return
	}
	user, err := n.userService.Login(input)
	if err != nil {
		log.Println(err)
		c.HTML(http.StatusOK, "login", gin.H{
			"title":   "Login SSL",
			"message": "Username/ Password yang anda masukkan Salah!",
		})
		return
	}
	session.Set("userID", user.User_id)
	session.Set("userName", user.Name)
	session.Save()

	c.Redirect(http.StatusFound, "/home")
}

func (h *userHandler) Logout(c *gin.Context) {
	session := sessions.Default(c)
	session.Clear()
	session.Save()

	http.SetCookie(c.Writer, &http.Cookie{
		Name:   "message",
		MaxAge: -1,
	})

	c.Redirect(http.StatusFound, "/login")
}
