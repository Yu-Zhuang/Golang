# Todo Web App 實作

## 技術簡介
1. 連接postgresql資料庫
2. JWT存放cookie驗證＆auth middleware&保持稱入狀態
3. 透過登入解析token去資料庫拿回對應資料
```go
// HomeHandler ...
func HomeHandler(c *gin.Context) {
	username := c.MustGet("username").(string)
	if username == "" {
		c.Redirect(http.StatusFound, "/login")
	} else {
		var fuser models.User
		var fTodo []models.Todo
		// 透過username從資料庫取得該user
		dbc.Db.Where("Username = ?", username).Find(&fuser)
		// 透過該user從資料庫提取他的todo list
		dbc.Db.Model(&fuser).Association("Todo").Find(&fTodo)
		// 回傳
		c.HTML(http.StatusOK, "home.html", gin.H{
			"content": username,
			"todo":    fTodo,
		})
	}
}
```

## 程式功能簡介
1. 註冊&登入&保持登入狀態
2. 新增代辦事項

## 使用框架
1. gin
2. gorm
3. jwt

## 其他
1. 預設資料庫為postgresql, "host=localhost user=postgres password=abc123 dbname=postgres port=1036 sslmode=disable TimeZone=Asia/Shanghai"
2. 預設server port: 8080