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

type BlogItem struct {
	title          string
	content        string
}

type AdminController struct{}

func (ac *AdminController) AddBlogCtr(c *gin.Context) {
	c.HTML(http.StatusOK, "add-blog.html", gin.H{})
}

func (ac *AdminController) SaveBlogAddCtr(c *gin.Context) {
	var BI BlogItem
	config := GetConfig()
	c.BindWith(&BI, binding.Form)
	if BI.title == "" {
		ShowMessage(c, "Title can not empty")
		return
	}
	if BI.content == "" {
		ShowMessage(c, "Content can not empty")
		return
	}
	db := GetDB(config)
	_, err := db.Exec("insert into top_article (title, content) values (?, ?)", BI.title, BI.content);
	if err == nil {
		ShowMessage(c, "Success")
	} else {
		ShowMessage(c, "Failed to save blog")
	}

}

func (ac *AdminController) LoginCtr(c *gin.Context) {
	c.HTML(http.StatusOK, "admin-login.html", gin.H{})
}

func (ac *AdminController) LoginProcessCtr(c *gin.Context) {
	var form AdminLoginForm
	config := GetConfig()
	c.BindWith(&form, binding.Form)

	if form.Username == config.Admin_user && form.Password == config.Admin_password {
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
