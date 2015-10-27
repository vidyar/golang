package main

import (
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/naoina/toml"
	"html/template"
	"io/ioutil"
	"net/http"
	"os"
)

/*
	Show message with template
*/
func ShowMessage(c *gin.Context, message string) {
	c.HTML(http.StatusOK, "message.html", gin.H{
		"message": template.HTML(message),
	})
}

func GetDB(config *appConfig) *sql.DB {
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%d)/gosense?charset=utf8mb4", config.Db_user, config.Db_password, config.Db_host, config.Db_port))
	if err != nil {
		panic(err.Error())
	}
	db.SetMaxIdleConns(50)
	db.SetMaxOpenConns(100)
	return db
}

func GetConfig() *appConfig {
	f, err := os.Open("config.toml")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	buf, err := ioutil.ReadAll(f)
	if err != nil {
		panic(err)
	}
	var config appConfig
	if err := toml.Unmarshal(buf, &config); err != nil {
		panic(err)
	}
	return &config
}
