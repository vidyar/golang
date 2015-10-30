package main

import (
	"database/sql"
	"fmt"
	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"html/template"
	"log"
	"net/http"
	"strconv"
"strings"
)

type FrontController struct{}

func (fc *FrontController) AboutCtr(c *gin.Context) {
	c.HTML(http.StatusOK, "about.html", gin.H{})
}
func (fc *FrontController) PingCtr(c *gin.Context) {
	c.String(http.StatusOK, "pong")
}
func (fc *FrontController) HomeCtr(c *gin.Context) {
	config := GetConfig()
	db := GetDB(config)
	defer db.Close()
	page, err := strconv.Atoi(c.DefaultQuery("page", "1"))
	if err != nil {
		log.Fatal(err)
	}
	page -= 1
	if page < 0 {
		page = 0
	}

	prev_page := page
	if prev_page < 1 {
		prev_page = 1
	}
	next_page := page + 2

	var blogList string
	rpp := 20
	offset := page * rpp
	log.Println(rpp)
	log.Println(offset)
	rows, err := db.Query("Select aid, title from top_article where publish_status = 1 order by aid desc limit ? offset ? ", &rpp, &offset)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	var (
		aid   int
		title sql.NullString
	)
	for rows.Next() {
		err := rows.Scan(&aid, &title)
		if err != nil {
			log.Fatal(err)
		}
		blogList += fmt.Sprintf(
			"<li><a href=\"/view/%d\">%s</a></li>",
			aid,
			title.String,
		)
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}
	session := sessions.Default(c)
	username := session.Get("username")
	c.HTML(http.StatusOK, "index.html", gin.H{
		"bloglist":  template.HTML(blogList),
		"username":  username,
		"prev_page": prev_page,
		"next_page": next_page,
	})
}

func (fc *FrontController) SearchCtr(c *gin.Context) {
	config := GetConfig()
	db := GetDB(config)
	defer db.Close()
	page, err := strconv.Atoi(c.DefaultQuery("page", "1"))
	if err != nil {
		log.Fatal(err)
	}
	page -= 1
	if page < 0 {
		page = 0
	}

	prev_page := page
	if prev_page < 1 {
		prev_page = 1
	}
	next_page := page + 2
	keyword := c.DefaultQuery("keyword", "")
	if keyword == "" {
		(&msg{"Keyword can not empty"}).ShowMessage(c)
		return
	}
	keyword = strings.Replace(keyword, " ", "%", -1)

	var blogList string
	rpp := 20
	offset := page * rpp
	rows, err := db.Query(
		"Select aid, title from top_article where publish_status = 1 and (title like ? or content like ?) order by aid desc limit ? offset ? ",
		"%"+keyword+"%", "%"+keyword+"%", &rpp, &offset)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	var (
		aid   int
		title sql.NullString
	)
	for rows.Next() {
		err := rows.Scan(&aid, &title)
		if err != nil {
			log.Fatal(err)
		}
		blogList += fmt.Sprintf(
			"<li><a href=\"/view/%d\">%s</a></li>",
			aid,
			title.String,
		)
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}
	session := sessions.Default(c)
	username := session.Get("username")
	c.HTML(http.StatusOK, "search.html", gin.H{
		"bloglist":  template.HTML(blogList),
		"keyword":   keyword,
		"username":  username,
		"prev_page": prev_page,
		"next_page": next_page,
	})
}

func (fc *FrontController) ViewAltCtr(c *gin.Context) {
	id := c.DefaultQuery("id", "0")
	c.Redirect(301, fmt.Sprintf("/view/%s", id))

}
func (fc *FrontController) ViewCtr(c *gin.Context) {
	id := c.Param("id")
	config := GetConfig()
	db := GetDB(config)
	defer db.Close()
	rows, err := db.Query("Select * from top_article where aid = ?", &id)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	var (
		aid            int
		title          sql.NullString
		content        sql.NullString
		publish_time   sql.NullString
		publish_status sql.NullInt64
	)
	for rows.Next() {
		err := rows.Scan(&aid, &title, &content, &publish_time, &publish_status)
		if err != nil {
			log.Fatal(err)
		}
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}
	c.HTML(http.StatusOK, "view.html", gin.H{"aid": aid, "title": title.String, "content": template.HTML(content.String), "publish_time": publish_time.String})

}
