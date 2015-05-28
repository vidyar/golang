package main

import (
	"github.com/gin-gonic/gin"
	"github.com/netroby/gosense/controller/admin"
	"github.com/netroby/gosense/controller/front"
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