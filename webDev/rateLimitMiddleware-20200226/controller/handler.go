package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// HomeHandler : home handler
func HomeHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"msg":      "hello Dcard",
		"describe": "rate limit middleware tester, check out inspect of your browser",
	})
}
