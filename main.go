package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.LoadHTMLGlob("templates/*")
	r.Static("/assets", "./assets")

	r.GET("/", FrontHomeCtr)
	r.GET("/ping", FrontPingCtr)

	v1 := r.Group("/admin")
	{
		v1.GET("/login", AdminLoginCtr)
		v1.POST("/login-process", AdminLoginProcessCtr)
		v1.GET("/addblog", AdminAddBlogCtr)
	}
	// Listen and serve on 0.0.0.0:8080
	r.Run(":8080")
}
