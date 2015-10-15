package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.LoadHTMLGlob("templates/*")
	r.Static("/assets", "./assets")

	fc := new(FrontController)
	r.GET("/", fc.HomeCtr)
	r.GET("/ping", fc.PingCtr)

	ac := new(AdminController)
	v1 := r.Group("/admin")
	{
		v1.GET("/login", ac.LoginCtr)
		v1.POST("/login-process", ac.LoginProcessCtr)
		v1.GET("/addblog", ac.AddBlogCtr)
	}
	// Listen and serve on 0.0.0.0:8080
	r.Run(":8080")
}
