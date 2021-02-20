package main

import (
	"fmt"
	"todoWebApp/controller"
	"todoWebApp/dbc"
	"todoWebApp/models"

	"github.com/gin-gonic/gin"
)

func main() {
	// 連接數據庫
	err := dbc.InitDB()
	if err != nil {
		fmt.Println(err.Error())
	}
	// 產生table
	dbc.Db.AutoMigrate(&models.User{})
	dbc.Db.AutoMigrate(&models.Todo{})

	// 註冊路由
	r := gin.Default()
	r.LoadHTMLGlob("templates/*")

	// 各頁路由
	r.GET("/", controller.JWTAuthMiddleware(), controller.HomeHandler)
	r.GET("/login", controller.LoginHandler)
	r.GET("/signup", controller.SignupHandler)

	// 註冊帳號
	r.POST("/signup", controller.SignupAdd)
	// 登入驗證
	r.POST("/login", controller.LoginAuth)

	// 新增todo
	r.POST("/todoAdd", controller.TodoAdd)

	r.Run()
}
