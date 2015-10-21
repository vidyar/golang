package main

import (
	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"net/http"
)

type AdminLoginForm struct {
	Username string `form:"username" binding:"required"`
	Password string `form:"password" binding:"required"`
}

type AdminController struct{}

func (ac *AdminController) AddBlogCtr(c *gin.Context) {
	c.HTML(http.StatusOK, "add-blog.html", gin.H{})
}

func (ac *AdminController) LoginCtr(c *gin.Context) {
	c.HTML(http.StatusOK, "admin-login.html", gin.H{})
}

func (ac *AdminController) LoginProcessCtr(c *gin.Context) {
	var form AdminLoginForm
	c.BindWith(&form, binding.Form)

	if form.Username == "netroby" && form.Password == "dy123456" {
		session := sessions.Default(c)
		session.Set("username", "netroby")
		session.Save()
		c.Redirect(301, "/")
	} else {
		message := "Login failed"
		ShowMessage(c, message)
	}
}

func (ac *AdminController) LogoutCtr(c *gin.Context) {
	session := sessions.Default(c)
	session.Delete("username")
	session.Save()
	c.Redirect(301, "/")
}
