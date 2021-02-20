package controller

import (
	"net/http"
	"todoWebApp/dbc"
	"todoWebApp/jwtp"
	"todoWebApp/models"

	"github.com/gin-gonic/gin"
)

// TodoAdd :POST, 新增該人代辦事項
func TodoAdd(c *gin.Context) {
	// 取出cookie
	cookie, err := c.Cookie("auth_token")
	if err != nil {
		c.Redirect(http.StatusOK, "/login")
		return
	}
	// 解析cookie
	mc, err := jwtp.ParseToken(cookie)
	if err != nil {
		c.Redirect(http.StatusFound, "/login")
		return
	}
	// 查找資料庫是否有該人
	var fuser models.User
	dbc.Db.Where("Username = ?", mc.Username).Find(&fuser)
	// 提取todo
	var nTodo models.Todo
	c.Bind(&nTodo)
	if nTodo.Content == "" {
		c.Redirect(http.StatusFound, "/")
		return
	}
	// 創建todo + 該人FK
	nTodo.UserRefer = fuser.ID
	// 存入資料庫
	if err := models.CreateTodo(&nTodo); err != nil {
		c.Redirect(http.StatusFound, "/")
	} else {
		c.Redirect(http.StatusFound, "/")
	}
}
