package main

import (
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	jwt "gopkg.in/dgrijalva/jwt-go.v3"
)

func main() {
	r := gin.Default()
	r.LoadHTMLGlob("templates/*")
	// 首頁: 透過middleware驗證cookie中是否有正確的token, 有則可進 無則到登入
	r.GET("/", JWTAuthMiddleware(), homeHandler)
	// 進登入頁面
	r.GET("/signin", func(c *gin.Context) {
		c.HTML(http.StatusOK, "sign.html", nil)
	})
	// 進註冊頁面
	r.GET("/signup", func(c *gin.Context) {
		c.HTML(http.StatusOK, "signup.html", nil)
	})
	// 登入請求: 如正確會用帳密產生token並存在使用者的cookie中,錯誤則到註冊頁面
	r.POST("/auth", authHandler)
	// 啟動server
	r.Run()
}

func homeHandler(c *gin.Context) {
	// 這裏會拿到解析token後的username (或是說帳號)
	username := c.MustGet("username").(string)
	fmt.Println("cookie", username)
	if username == "" {
		c.Redirect(http.StatusFound, "/signin")
	} else {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"content": username,
		})
	}
}

func authHandler(c *gin.Context) {
	// 用户发送用户名和密码过来
	var user UserInfo
	err := c.ShouldBind(&user)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 2001,
			"msg":  "无效的参数",
		})
		return
	}
	// 校验用户名和密码是否正确
	if user.UserName == "john" && user.Password == "john123" {
		// 生成Token
		tokenString, _ := GenToken(user.UserName)
		c.SetCookie("auth_token", tokenString, 60, "/", "127.0.0.1", false, true)
		c.Redirect(http.StatusFound, "/")
	} else {
		c.Redirect(http.StatusFound, "/signup")
	}
	return
}

// JWTAuthMiddleware 基于JWT的认证中间件
func JWTAuthMiddleware() func(c *gin.Context) {
	return func(c *gin.Context) {
		// 客户端携带Token有三种方式 1.放在请求头 2.放在请求体 3.放在URI
		// 這裏取放在cookie中的
		authHeader, err := c.Cookie("auth_token")
		if err != nil {
			c.Redirect(http.StatusFound, "/signin")
			c.Abort()
			return
		}

		// 取到的tokenString，我们使用之前定义好的解析JWT的函数来解析它
		mc, err := ParseToken(authHeader)
		if err != nil {
			c.Redirect(http.StatusFound, "/signin")
			c.Abort()
			return
		}
		// 将当前请求的username信息保存到请求的上下文c上
		c.Set("username", mc.Username)
		c.Next() // 后续的处理函数可以用过c.Get("username")来获取当前请求的用户信息

	}
}

// MyClaims ...
type MyClaims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

// UserInfo ...
type UserInfo struct {
	UserName string `json:"userName" form:"userName"`
	Password string `json:"password" form:"password"`
}

// TokenExpireDuration ..
const TokenExpireDuration = time.Hour * 2

// MySecret 密鑰
var MySecret = []byte("我是密鑰")

// GenToken 生成JWT
func GenToken(username string) (string, error) {
	// 创建一个我们自己的声明
	c := MyClaims{
		username, // 自定义字段
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(TokenExpireDuration).Unix(), // 过期时间
			Issuer:    "my-project",                               // 签发人
		},
	}
	// 使用指定的签名方法创建签名对象
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, c)
	// 使用指定的secret签名并获得完整的编码后的字符串token
	return token.SignedString(MySecret)
}

// ParseToken 解析JWT
func ParseToken(tokenString string) (*MyClaims, error) {
	// 解析token
	token, err := jwt.ParseWithClaims(tokenString, &MyClaims{}, func(token *jwt.Token) (i interface{}, err error) {
		return MySecret, nil
	})
	if err != nil {
		return nil, err
	}
	if claims, ok := token.Claims.(*MyClaims); ok && token.Valid { // 校验token
		return claims, nil
	}
	return nil, errors.New("invalid token")
}
