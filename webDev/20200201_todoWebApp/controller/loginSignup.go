package controller

import (
	"fmt"
	"net/http"
	"todoWebApp/dbc"
	"todoWebApp/jwtp"
	"todoWebApp/models"

	"github.com/gin-gonic/gin"
)

// HomeHandler ...
func HomeHandler(c *gin.Context) {
	username := c.MustGet("username").(string)
	if username == "" {
		c.Redirect(http.StatusFound, "/login")
	} else {
		var fuser models.User
		var fTodo []models.Todo
		dbc.Db.Where("Username = ?", username).Find(&fuser)
		dbc.Db.Model(&fuser).Association("Todo").Find(&fTodo)
		c.HTML(http.StatusOK, "home.html", gin.H{
			"content": username,
			"todo":    fTodo,
		})
	}
}

// LoginHandler ...
func LoginHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "login.html", nil)
}

// SignupHandler ...
func SignupHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "signup.html", nil)
}

// SignupAdd ... 註冊
func SignupAdd(c *gin.Context) {
	// 提取新註冊帳號
	var newUser models.User
	c.Bind(&newUser)
	// 確認帳號是否重複註冊
	var fuser models.User
	dbc.Db.Where("Username = ?", newUser.Username).Find(&fuser)
	// 如已有此帳號
	if fuser.Username == newUser.Username {
		fmt.Println("重複帳號")
		c.Redirect(http.StatusFound, "/signup")
		return
	}
	// 如無 則創建新帳號
	if err := models.CreateUser(&newUser); err != nil {
		c.Redirect(http.StatusFound, "/signup")
	} else {
		c.Redirect(http.StatusFound, "/login")
	}
}

// LoginAuth 登入驗證
func LoginAuth(c *gin.Context) {
	// 用户发送用户名和密码过来
	var user models.User
	err := c.ShouldBind(&user)
	if err != nil {
		c.Redirect(http.StatusFound, "/signip")
		return
	}
	// 校验用户名和密码是否正确
	var fuser models.User
	dbc.Db.Where("Username = ? AND Password = ?", user.Username, user.Password).Find(&fuser)
	if fuser.Username == "" || user.Password == "" {
		fmt.Println("not found")
		c.Redirect(http.StatusFound, "/signup")
	} else {
		tokenString, _ := jwtp.GenToken(user.Username)
		c.SetCookie("auth_token", tokenString, 60, "/", "127.0.0.1", false, true)
		c.Redirect(http.StatusFound, "/")
	}
	return
}

// JWTAuthMiddleware 基于JWT的认证中间件
func JWTAuthMiddleware() func(c *gin.Context) {
	return func(c *gin.Context) {
		// 取放在cookie中的
		authHeader, err := c.Cookie("auth_token")
		if err != nil {
			c.Redirect(http.StatusFound, "/login")
			c.Abort()
			return
		}
		// 取到的tokenString，我们使用之前定义好的解析JWT的函数来解析它
		mc, err := jwtp.ParseToken(authHeader)
		if err != nil {
			c.Redirect(http.StatusFound, "/login")
			c.Abort()
			return
		}
		// 将当前请求的username信息保存到请求的上下文c上
		c.Set("username", mc.Username)
		c.Next() // 后续的处理函数可以用过c.Get("username")来获取当前请求的用户信息

	}
}
