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

type FrontController struct {
}

func (fc *FrontController) AboutCtr(c *gin.Context) {
	c.HTML(http.StatusOK, "about.html", gin.H{})
}
func (fc *FrontController) PingCtr(c *gin.Context) {
	c.String(http.StatusOK, "pong")
}
func (fc *FrontController) HomeCtr(c *gin.Context) {
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

	rpp := 20
	offset := page * rpp
	CKey := fmt.Sprintf("home-page-%d-rpp-%d", page, rpp)
	var blogList string
	val, ok := Cache.Get(CKey)
	if val != nil && ok == true {
		fmt.Println("Ok, we found cache, Cache Len: ", Cache.Len())
		blogList = val.(string)
	} else {
		rows, err := DB.Query("Select aid, title from top_article where publish_status = 1 order by aid desc limit ? offset ? ", &rpp, &offset)
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
		Cache.Add(CKey, blogList)
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
	fmt.Println(keyword)
	if len(keyword) <= 0 {
		(&msg{"Keyword can not empty"}).ShowMessage(c)
		return
	}
	orig_keyword := keyword
	keyword = strings.Replace(keyword, " ", "%", -1)

	var blogList string
	rpp := 20
	offset := page * rpp
	rows, err := DB.Query(
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
		"keyword":   orig_keyword,
		"username":  username,
		"prev_page": prev_page,
		"next_page": next_page,
	})
}

func (fc *FrontController) ViewAltCtr(c *gin.Context) {
	id := c.DefaultQuery("id", "0")
	c.Redirect(301, fmt.Sprintf("/view/%s", id))

}

type VBlogItem struct {
	aid            int
	title          sql.NullString
	content        sql.NullString
	publish_time   sql.NullString
	publish_status sql.NullInt64
}

func (fc *FrontController) ViewCtr(c *gin.Context) {
	id := c.Param("id")
	var blog VBlogItem
	CKey := fmt.Sprintf("blogitem-%d", id)
	val, ok := Cache.Get(CKey)
	if val != nil && ok == true {
		fmt.Println("Ok, we found cache, Cache Len: ", Cache.Len())
		blog = val.(VBlogItem)
	} else {
		rows, err := DB.Query("Select * from top_article where aid = ?", &id)
		if err != nil {
			log.Fatal(err)
		}
		defer rows.Close()
		var ()
		for rows.Next() {
			err := rows.Scan(&blog.aid, &blog.title, &blog.content, &blog.publish_time, &blog.publish_status)
			if err != nil {
				log.Fatal(err)
			}
		}
		err = rows.Err()
		if err != nil {
			log.Fatal(err)
		}
		Cache.Add(CKey, blog)
	}
	c.HTML(http.StatusOK, "view.html", gin.H{
		"aid":          blog.aid,
		"title":        blog.title.String,
		"content":      template.HTML(blog.content.String),
		"publish_time": blog.publish_time.String,
	})

}
