package main

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"net/http"
	"strconv"
	"log"
	"fmt"
"html/template"
)

type AdminLoginForm struct {
	Username string `form:"username" binding:"required"`
	Password string `form:"password" binding:"required"`
}

type BlogItem struct {
	Title   string `form:"title" binding:"required"`
	Content string `form:"content" binding:"required"`
}

type AdminController struct{}

func (ac *AdminController) ListBlogCtr(c *gin.Context) {
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
			"<li><a href=\"/view/%d\">%s</a>    [<a href=\"/admin/editblog/%d\">Edit</a>] [<a href=\"/admin/deleteblog/%d\">Delete</a>]</li>",
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
	c.HTML(http.StatusOK, "admin.list.blog.html", gin.H{
		"bloglist":  template.HTML(blogList),
		"username":  username,
		"prev_page": prev_page,
		"next_page": next_page,
	})
}

func (ac *AdminController) EditBlogCtr(c *gin.Context) {
	ShowMessage(c, "This is EditBlog action")
}
func (ac *AdminController) DeleteBlogCtr(c *gin.Context) {
	ShowMessage(c, "This is delete blog action")
}


func (ac *AdminController) AddBlogCtr(c *gin.Context) {
	c.HTML(http.StatusOK, "add-blog.html", gin.H{})
}

func (ac *AdminController) SaveBlogAddCtr(c *gin.Context) {
	var BI BlogItem
	config := GetConfig()
	c.BindWith(&BI, binding.Form)
	if BI.Title == "" {
		ShowMessage(c, "Title can not empty")
		return
	}
	if BI.Content == "" {
		ShowMessage(c, "Content can not empty")
		return
	}
	db := GetDB(config)
	_, err := db.Exec("insert into top_article (title, content) values (?, ?)", BI.Title, BI.Content)
	if err == nil {
		ShowMessage(c, "Success")
	} else {
		ShowMessage(c, "Failed to save blog")
	}

}

func (ac *AdminController) LoginCtr(c *gin.Context) {
	c.HTML(http.StatusOK, "admin-login.html", gin.H{})
}

func (ac *AdminController) LoginProcessCtr(c *gin.Context) {
	var form AdminLoginForm
	config := GetConfig()
	c.BindWith(&form, binding.Form)

	if form.Username == config.Admin_user && form.Password == config.Admin_password {
		session := sessions.Default(c)
		session.Set("username", "netroby")
		session.Save()
		c.Redirect(301, "/")
	} else {
		message := "Login failed"
		ShowMessage(c, message)
	}
}

func (ac *AdminController) LogoutCtr(c *gin.Context) {
	session := sessions.Default(c)
	session.Delete("username")
	session.Save()
	c.Redirect(301, "/")
}
