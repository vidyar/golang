package admin

import (
	"time"
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
		expire := time.Now().AddDate(0, 0, 1)
		cookie := http.Cookie{Name: "username", Value: "netroby", Path: "/", Expires: expire, MaxAge: 86400}

		http.SetCookie(c.Writer, &cookie)
		c.String(http.StatusOK, "you are logged in<a href=\"/\">Click to go</a>")
	} else {
		c.String(http.StatusOK, "Login failed" + form.Username)
	}
}