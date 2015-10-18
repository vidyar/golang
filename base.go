package main

import (
	"github.com/gin-gonic/gin"
	"html/template"
	"net/http"
)

/*
	Show message with template
*/
func ShowMessage(c *gin.Context, message string) {
	c.HTML(http.StatusOK, "message.html", gin.H{
		"message": template.HTML(message),
	})
}
