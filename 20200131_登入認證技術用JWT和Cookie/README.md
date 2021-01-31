# 登入驗證使用JWT和Cookie

## 使用
1. 下載包: git clone 網址
2. 進入該資料後編譯main.go: go build
3. 啟動編譯後檔案
4. 開瀏覽器輸入網址: 127.0.0.1:8080

## 實作程式碼簡介
0. 主程式
```go
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
```
1. 如何產生與解析token: 此程式碼是利用別人寫好的genToken和parseToken
2. 如何將token存到使用者瀏覽器並保持登入
```go
// 比對登入的帳密是否正確
if user.UserName == "john" && user.Password == "john123" {
	// 正確生成Token
	tokenString, _ := GenToken(user.UserName)
	// 將生成的token儲存到瀏覽器中, 在此設定保存時間為60秒, 表示可保持登入狀態60秒
	c.SetCookie("auth_token", tokenString, 60, "/", "127.0.0.1", false, true)
	c.Redirect(http.StatusFound, "/")
} else {
	c.Redirect(http.StatusFound, "/signup")
}
```
3. 如何知道該token是誰
```go
// 取到的tokenString，我们使用之前定义好的解析JWT的函数来解析它
mc, err := ParseToken(authHeader)
if err != nil {
	c.Redirect(http.StatusFound, "/signin") //錯誤的token則切到登入頁面
	c.Abort()
	return
}
// 正確的話, 由於token是由帳號加密而成, 因此解析後可得帳號(username)
c.Set("username", mc.Username) // 將请求的username信息保存到请求的上下文c上
c.Next() // 后续的处理函数可以用过c.Get("username")来获取当前请求的用户信息
```
---

by鈺粧
2020.01.31


