package user

import (
	"fmt"
	"log"
	"net/http"
	// "strings"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/mhilmi999/Tugas-Akhir-Smart-Street-Light-LoRaWAN/modules/v1/utilities/user/models"
	"github.com/mhilmi999/Tugas-Akhir-Smart-Street-Light-LoRaWAN/modules/v1/utilities/user/service"
	// helper "github.com/mhilmi999/Tugas-Akhir-Smart-Street-Light-LoRaWAN/pkg/helpers"
)

type UserHandler interface {
	Register(c *gin.Context)
	Login(c *gin.Context)
	ReceivedData(c *gin.Context)
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

func (n *userHandler) Logout(c *gin.Context) {
	session := sessions.Default(c)
	session.Clear()
	session.Save()

	http.SetCookie(c.Writer, &http.Cookie{
		Name:   "message",
		MaxAge: -1,
	})

	c.Redirect(http.StatusFound, "/login")
}

func (n *userHandler) ReceivedData(c *gin.Context){
	// var input models.Received
	// if err := c.ShouldBindJSON(&input); err != nil{
	// 	response := helper.APIRespon("Error inputnya invalid", 220, "error", nil)
	// 	c.JSON(220,response)
	// 	return
	// }
	// Antares_Id := strings.Replace(input.First.Pi, "/antares-cse/cnt-", "", -1)
	// fmt.Println(input)
	// id := n.userService.GetId(Antares_Id)
	// getData := n.userService.GetData(input.First.Con)
	// inputData := n.userService.DataCheck(id, getData)
	// if inputData == 2 {
	// 	response := helper.APIRespon("Error, data headers ga sesuai. ", 210, "error", nil)
	// 	c.JSON(210, response)
	// } else if inputData == 1 {
	// 	response := helper.APIRespon("Success to input data", http.StatusOK, "success", input)
	// 	c.JSON(http.StatusOK, response)
	// } else { //0 gagal
	// 	response := helper.APIRespon("Error, unable to enter data", 215, "error", nil)
	// 	c.JSON(215, response)
	// }

	token := "01fe7c50a39803d0:93a1cf61893c1605"
	getLatestData, err := n.userService.GetLatestData(token)
	if err !=nil{
		fmt.Println(err)
		return
	}
	fmt.Println(getLatestData)
}