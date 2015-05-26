package front


import (
	"fmt"
	"html/template"
	"net/http"
	"github.com/gin-gonic/gin"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

type blogItem struct {
	url   string
	title string
}

func PingCtr(c *gin.Context) {
	c.String(http.StatusOK, "pong")
}
func HomeCtr(c *gin.Context) {
	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/ultrax")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err.Error())
	}
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