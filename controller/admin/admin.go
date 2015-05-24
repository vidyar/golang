package admin

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

type LoginForm struct {
	Username     string `form:"username" binding:"required"`
	Password string `form:"password" binding:"required"`
}


func LoginCtr(c *gin.Context) {
	c.HTML(http.StatusOK, "admin-login.html", gin.H{})
}

func LoginProcessCtr(c *gin.Context) {
	var form LoginForm
	c.BindWith(&form, binding.MultipartForm)

	if form.Username == "netroby" && form.Password == "deyilife" {
		c.String(http.StatusOK, "you are logged in")
	} else {
		c.String(http.StatusOK, "Login failed" + form.Username)
	}
}