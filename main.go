package main

import (
	"./internal/controller/admin"
	"./internal/controller/front"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.LoadHTMLGlob("templates/*")
	r.Static("/assets", "./assets")

	r.GET("/", front.HomeCtr)
	r.GET("/ping", front.PingCtr)

	v1 := r.Group("/admin")
	{
		v1.GET("/login", admin.LoginCtr)
		v1.POST("/login-process", admin.LoginProcessCtr)
		v1.GET("/addblog", admin.AddBlogCtr)
	}
	// Listen and serve on 0.0.0.0:8080
	r.Run(":8080")
}
