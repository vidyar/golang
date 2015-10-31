package main

import (
	"database/sql"
	"fmt"
	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	_ "github.com/go-sql-driver/mysql"
	"html/template"
	"log"
	"net/http"
	"strconv"
	"time"
)

type AdminLoginForm struct {
	Username string `form:"username" binding:"required"`
	Password string `form:"password" binding:"required"`
}

type BlogItem struct {
	Title   string `form:"title" binding:"required"`
	Content string `form:"content" binding:"required"`
}
type EditBlogItem struct {
	Aid     string `form:"aid" binding:"required"`
	Title   string `form:"title" binding:"required"`
	Content string `form:"content" binding:"required"`
}

type AdminController struct{
}

func (ac *AdminController) ListBlogCtr(c *gin.Context) {

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
	(&msg{"This is EditBlog action"}).ShowMessage(c)
}

func (ac *AdminController) DeleteBlogCtr(c *gin.Context) {
	(&msg{"This is delete blog action"}).ShowMessage(c)
}

func (ac *AdminController) AddBlogCtr(c *gin.Context) {
	c.HTML(http.StatusOK, "add-blog.html", gin.H{})
}

func (ac *AdminController) SaveBlogEditCtr(c *gin.Context) {
	var BI EditBlogItem
	c.BindWith(&BI, binding.Form)
	if BI.Aid == "" {
		(&msg{"Can not find the blog been edit"}).ShowMessage(c)
		return
	}
	if BI.Title == "" {
		(&msg{"Title can not empty"}).ShowMessage(c)
		return
	}
	if BI.Content == "" {
		(&msg{"Content can not empty"}).ShowMessage(c)
		return
	}
	_, err := DB.Exec("update top_article set title=?, content=? where aid = ?", BI.Title, BI.Content, BI.Aid)
	if err == nil {
		(&msg{"Success"}).ShowMessage(c)
	} else {
		(&msg{"Failed to save blog"}).ShowMessage(c)
	}

}
func (ac *AdminController) SaveBlogAddCtr(c *gin.Context) {
	var BI BlogItem
	c.BindWith(&BI, binding.Form)
	if BI.Title == "" {
		(&msg{"Title can not empty"}).ShowMessage(c)
		return
	}
	if BI.Content == "" {
		(&msg{"Content can not empty"}).ShowMessage(c)
		return
	}
	_, err := DB.Exec(
		"insert into top_article (title, content, publish_time, publish_status) values (?, ?, ?, 1)",
		BI.Title, BI.Content, time.Now().Format("2006-01-02 15:04:05"))
	if err == nil {
		(&msg{"Success"}).ShowMessage(c)
	} else {
		(&msg{"Failed to save blog"}).ShowMessage(c)
	}

}

func (ac *AdminController) LoginCtr(c *gin.Context) {
	c.HTML(http.StatusOK, "admin-login.html", gin.H{})
}

func (ac *AdminController) LoginProcessCtr(c *gin.Context) {
	var form AdminLoginForm
	c.BindWith(&form, binding.Form)

	if form.Username == Config.Admin_user && form.Password == Config.Admin_password {
		session := sessions.Default(c)
		session.Set("username", "netroby")
		session.Save()
		c.Redirect(301, "/")
	} else {
		(&msg{"Login failed"}).ShowMessage(c)
	}
}

func (ac *AdminController) LogoutCtr(c *gin.Context) {
	session := sessions.Default(c)
	session.Delete("username")
	session.Save()
	c.Redirect(301, "/")
}
