package main

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"html/template"
	"net/http"
	"time"
)

type AdminLoginForm struct {
	Username string `form:"username" binding:"required"`
	Password string `form:"password" binding:"required"`
}

func AdminAddBlogCtr(c *gin.Context) {
	c.HTML(http.StatusOK, "add-blog.html", gin.H{})
}

func AdminLoginCtr(c *gin.Context) {
	c.HTML(http.StatusOK, "admin-login.html", gin.H{})
}

func AdminLoginProcessCtr(c *gin.Context) {
	var form AdminLoginForm
	c.BindWith(&form, binding.Form)

	if form.Username == "netroby" && form.Password == "deyilife" {
		expire := time.Now().AddDate(0, 0, 1)
		cookie := http.Cookie{Name: "username", Value: "netroby", Path: "/", Expires: expire, MaxAge: 86400}

		http.SetCookie(c.Writer, &cookie)
		message := "you are logged in<a href=\"/\">Click to go</a>"
		AdminShowMessage(c, message)
	} else {
		message := "Login failed"
		AdminShowMessage(c, message)
	}
}

/*
	Show message with template
*/
func AdminShowMessage(c *gin.Context, message string) {
	c.HTML(http.StatusOK, "message.html", gin.H{
		"message": template.HTML(message),
	})
}
