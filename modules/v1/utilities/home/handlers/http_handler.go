package home

import (
	"fmt"
	"strings"

	"github.com/gin-gonic/gin"
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
