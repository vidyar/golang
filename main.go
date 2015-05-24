package main

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)
func pingCtr(c *gin.Context) {
	c.String(http.StatusOK, "pong")
}
func homeCtr(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", gin.H{})
}
func loginCtr(c *gin.Context) {
	c.HTML(http.StatusOK, "admin-login.html", gin.H{})
}
func loginProcessCtr(c *gin.Context) {
	var form LoginForm
	c.BindWith(&form, binding.MultipartForm)

	if form.Username == "netroby" && form.Password == "deyilife" {
		c.String(http.StatusOK, "you are logged in")
	} else {
		c.String(http.StatusOK, "Login failed" + form.Username)
	}
}
type LoginForm struct {
	Username     string `form:"username" binding:"required"`
	Password string `form:"password" binding:"required"`
}

func main() {
	r := gin.Default()
	r.LoadHTMLGlob("templates/*")
	r.Static("/assets", "./assets")
	r.GET("/", homeCtr)
	r.GET("/ping", pingCtr)

	v1 := r.Group("/admin")
	{
		v1.GET("/login", loginCtr)
		v1.POST("/login-process", loginProcessCtr)
	}
	// Listen and serve on 0.0.0.0:8080
	r.Run(":8080")
}