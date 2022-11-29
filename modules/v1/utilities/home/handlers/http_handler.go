package home

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func (n *homeHandler) ReceivedData(c *gin.Context) {
	token := "01fe7c50a39803d0:93a1cf61893c1605"
	getLatestCon, err := n.homeService.GetLatestCon(token)
	if err != nil {
		fmt.Println(err)
		return
	}
	getDataCon, err := n.homeService.GetDatafromCon(getLatestCon.First.Con)
	if err != nil{
		fmt.Println(err)
		return
	}
	fmt.Println("Ini di handler \n",getDataCon)
}
