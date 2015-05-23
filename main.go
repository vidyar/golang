package main

import (
	"net/http"
	"github.com/gin-gonic/gin"
)
func pingCtr(c *gin.Context) {
	c.String(http.StatusOK, "pong")
}
func homeCtr(c *gin.Context) {
	c.String(http.StatusOK, "home control")
}

func main() {
	r := gin.Default()
	r.Static("/assets", "./assets")
	r.GET("/", homeCtr)
	r.GET("/ping", pingCtr)

	// Listen and serve on 0.0.0.0:8080
	r.Run(":8080")
}