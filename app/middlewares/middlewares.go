package middlewares

import (
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func IsLogin() gin.HandlerFunc{
	return func (c *gin.Context)  {
		session := sessions.Default(c)
		userIDSess := session.Get("userID")
		if userIDSess == nil{
			c.Redirect(http.StatusFound, "/login")
			return
		}
	}
}