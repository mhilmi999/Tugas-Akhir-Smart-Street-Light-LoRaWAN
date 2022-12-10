package home

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/mhilmi999/Tugas-Akhir-Smart-Street-Light-LoRaWAN/modules/v1/utilities/home/models"
	"github.com/mhilmi999/Tugas-Akhir-Smart-Street-Light-LoRaWAN/pkg/helpers"
)

func (n *homeHandler) ReceivedData(c *gin.Context) {
	token := "01fe7c50a39803d0:93a1cf61893c1605"
	getLatestCon, err := n.homeService.GetLatestCon(token)
	if err != nil {
		fmt.Println(err)
		return
	}
	Antares_Device_Id := strings.Replace(getLatestCon.First.Pi, "/antares-cse/cnt-", "", -1)
	fmt.Println("ini Antares Device ID nya \n",Antares_Device_Id)
	_, err = n.homeService.GetDatafromCon(getLatestCon.First.Con, Antares_Device_Id)
	if err != nil{
		fmt.Println(err)
		return
	}
	fmt.Println("Sukses masuk data ke db")
}

func (n *homeHandler) SubscribeWebhook(c *gin.Context) {
	var webhookData models.ObjectAntares1
	if err := c.ShouldBindJSON(&webhookData); err != nil {
		response := helpers.APIRespon("Error, inputan tidak sesuai", 220, "error",nil)
		c.JSON(220, response)
		return
	}
	fmt.Println("Pakettt \n",webhookData)
	Antares_Device_Id := strings.Replace(webhookData.First.M2m_nev.M2m_rep.M2m_cin.Pi, "/antares-cse/cnt-", "", -1)
	getData, err := n.homeService.GetDatafromWebhook(webhookData.First.M2m_nev.M2m_rep.M2m_cin.Con, Antares_Device_Id)
	if err != nil{
		fmt.Println(err)
		return
	}
	fmt.Println("Ini hasil data dari webhook \n",getData)
}

func (n *homeHandler) ControlLight(c *gin.Context){
	sessions := sessions.Default(c)
	power := c.Param("power")
	deviceid := c.Param("deviceid")
	token := "01fe7c50a39803d0:93a1cf61893c1605"

	_ = n.homeService.ControlLight(power, deviceid, token)
	if power == "1"{
		sessions.AddFlash("Device Menyala")
	}else if power == "0"{
		sessions.AddFlash("Device Mati")
	}else{
		sessions.AddFlash("Kesalahan! Tidak ada respon dari device")
	}
	c.Redirect(http.StatusFound, "/list-device")

}

func (n *homeHandler) SubscribeMqtt(c *gin.Context) {
	// n.homeService.SubscribeMqtt()
}