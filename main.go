package main

import (
	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
)

func main() {

	r := gin.Default()
	r.Static("/assets", "assets")
	store := sessions.NewCookieStore([]byte("gssecret"))
	r.Use(sessions.Sessions("mysession", store))
	r.LoadHTMLGlob("templates/*")

	fc := new(FrontController)
	r.GET("/", fc.HomeCtr)
	r.GET("/about", fc.AboutCtr)
	r.GET("/view/:id", fc.ViewCtr)
	r.GET("/view.php", fc.ViewAltCtr)
	r.GET("/ping", fc.PingCtr)
	r.GET("/search", fc.SearchCtr)

	ac := new(AdminController)
	v1 := r.Group("/admin")
	{
		v1.GET("/login", ac.LoginCtr)
		v1.POST("/login-process", ac.LoginProcessCtr)
		v1.GET("/logout", ac.LogoutCtr)
		v1.GET("/addblog", ac.AddBlogCtr)
		v1.POST("/save-blog-add", ac.SaveBlogAddCtr)
		v1.GET("/listblog", ac.ListBlogCtr)
		v1.GET("/deleteblog/:id", ac.DeleteBlogCtr)
		v1.GET("/editblog/:id", ac.EditBlogCtr)
	}
	// Listen and serve on 0.0.0.0:8080
	r.Run(":8080")
}
