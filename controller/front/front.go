package front


import (
	"net/http"
	"github.com/gin-gonic/gin"
)


func PingCtr(c *gin.Context) {
	c.String(http.StatusOK, "pong")
}
func HomeCtr(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", gin.H{})
}