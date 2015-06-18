package admin

import (
	"time"
	"net/http"
	"html/template"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

type LoginForm struct {
	Username string `form:"username" binding:"required"`
	Password string `form:"password" binding:"required"`
}

func AddBlogCtr(c *gin.Context) {
	c.HTML(http.StatusOK, "add-blog.html", gin.H{})
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
		message := "you are logged in<a href=\"/\">Click to go</a>"
		showMessage(c, message)
	} else {
		message := "Login failed"
		showMessage(c, message)
	}
}
/*
	Show message with template
 */
func showMessage(c *gin.Context, message string) {
	c.HTML(http.StatusOK, "message.html", gin.H{
		"message": template.HTML(message),
	})
}