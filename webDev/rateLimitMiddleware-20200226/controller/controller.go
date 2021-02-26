package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
	"rateLimitMiddleware/conf"
	"rateLimitMiddleware/dao"
	"rateLimitMiddleware/models"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

// HomeHandler : home handler
func HomeHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"msg": "hello rate limit middleware tester",
	})
}

// RateLimitMiddleware : rate limit
func RateLimitMiddleware(c *gin.Context) {
	var user models.RateLimit
	user.IP = c.Request.RemoteAddr
	// check weather database have user record
	value := dao.DB.Get(user.IP)
	fmt.Println(value.Val())
	// if has
	if value.Err() == nil {
		_ = json.Unmarshal([]byte(value.Val()), &user.RateLimitValue)
		fmt.Println(user.RateLimitValue)
		// 如果過期
		if time.Now().After(user.RateLimitValue.ExpireTime) == true {
			fmt.Println("4")
			// 創建新的並寫入db 與 res header
			nUser := CreateNewUserRateLimit()
			b, _ := json.Marshal(&(nUser.RateLimitValue))
			dao.DB.Set(user.IP, string(b), conf.RateLimitDuration)
			WriteRateLimitHeader(c, strconv.Itoa(nUser.RemainNum), nUser.ExpireTime.String())
			c.Next()
			return
		}
		// 未過期 則 取出剩下次數 -=1
		user.RateLimitValue.RemainNum--
		WriteRateLimitHeader(c, strconv.Itoa(user.RateLimitValue.RemainNum), user.RateLimitValue.ExpireTime.String())
		// if > rate limit
		if user.RateLimitValue.RemainNum <= 0 {
			fmt.Println("3")
			// return 429
			c.JSON(http.StatusTooManyRequests, gin.H{
				"msg": "too many request",
			})
			c.Abort()
			return
		}
		fmt.Println("2")
		b, _ := json.Marshal(&(user.RateLimitValue))
		fmt.Println(string(b))
		// 未超出 save to db
		dao.DB.Set(user.IP, string(b), user.RateLimitValue.ExpireTime.Sub(time.Now()))
		// next()
		c.Next()
		return
	}

	// not has
	nUser := CreateNewUserRateLimit()
	fmt.Println("1", nUser.RateLimitValue)
	b, _ := json.Marshal(&(nUser.RateLimitValue))
	fmt.Println(string(b))
	dao.DB.Set(user.IP, string(b), conf.RateLimitDuration)
	// writer response header
	WriteRateLimitHeader(c, strconv.Itoa(nUser.RemainNum), nUser.ExpireTime.String())
	// next()
	c.Next()
}

// CreateNewUserRateLimit ...
func CreateNewUserRateLimit() models.RateLimit {
	var user models.RateLimit
	user.RateLimitValue.ExpireTime = time.Now().Add(conf.RateLimitDuration)
	user.RateLimitValue.RemainNum = conf.RateLimitNum
	return user
}

// WriteRateLimitHeader ...
func WriteRateLimitHeader(c *gin.Context, remaining string, reset string) {
	c.Writer.Header().Set("X-RateLimit-Remaining", remaining)
	c.Writer.Header().Set("X-RateLimit-Reset", reset)
}
