package main

import (
	"database/sql"
	"fmt"
	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"html/template"
	"net/http"
	"log"
)

type FrontController struct{}

func getDB() (*sql.DB) {
	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/gosense?charset=utf8mb4")
	if err != nil {
		panic(err.Error())
	}
	db.SetMaxIdleConns(50)
	db.SetMaxOpenConns(100)
	return db
}

func (fc *FrontController) PingCtr(c *gin.Context) {
	c.String(http.StatusOK, "pong")
}
func (fc *FrontController) HomeCtr(c *gin.Context) {
	db := getDB()
	defer db.Close()
	var blogList string
	rpp := 20
	offset := 0
	rows, err := db.Query("Select aid, title from top_article order by aid desc limit ? offset ? ", &rpp, &offset)
	if err != nil {
		log.Fatal(err)
	}
	defer  rows.Close()
	var (
		aid int
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
		"bloglist": template.HTML(blogList),
		"username": username,
	})
}
func (fc *FrontController) ViewCtr(c *gin.Context) {
	id := c.Param("id")
	db := getDB()
	defer db.Close()
	rows, err := db.Query("Select * from top_article where aid = ?", &id)
	if err != nil {
		log.Fatal(err)
	}
	defer  rows.Close()
	var (
		aid int
		title sql.NullString
		content sql.NullString
		publish_time sql.NullString
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
	c.HTML(http.StatusOK, "view.html", gin.H{"aid": aid, "title": title, "content": content, "publish_time": publish_time})

}
