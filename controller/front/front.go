package front


import (
	"fmt"
	"html/template"
	"net/http"
	"github.com/gin-gonic/gin"
)

type blogItem struct {
	url   string
	title string
}

func PingCtr(c *gin.Context) {
	c.String(http.StatusOK, "pong")
}
func HomeCtr(c *gin.Context) {
	var bl [2]blogItem
	bl[0] =  blogItem{
		"//www.netroby.com/view.php?id=3833",
		"How To Manually Install Oracle Java on a Debian or Ubuntu VPS",
	}
	bl[1] = blogItem{
		"//www.netroby.com/view.php?id=3832",
		"Linux 4.0 kernel released",
	}
	var blogList string
	for i := 0; i < len(bl); i++ {
		blogList += fmt.Sprintf(
			"<li><a href=\"%s\">%s</a></li>",
			bl[i].url,
			bl[i].title,
		)
	}

	c.HTML(http.StatusOK, "index.html", gin.H{"bloglist":template.HTML(blogList)})
}