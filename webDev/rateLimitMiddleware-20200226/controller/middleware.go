package controller

import (
	"encoding/json"
	"net/http"
	"rateLimitMiddleware/conf"
	"rateLimitMiddleware/dao"
	"rateLimitMiddleware/logic"
	"rateLimitMiddleware/models"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

// RateLimitMiddleware : rate limi
func RateLimitMiddleware(c *gin.Context) {
	var user models.RateLimit
	user.IP = c.Request.RemoteAddr
	// check weather database have user record
	value := dao.DB.Get(user.IP)
	// if has user record or not expire
	if value.Err() == nil {
		json.Unmarshal([]byte(value.Val()), &user.RateLimitValue)
		user.RemainNum--
		// if > rate limit
		if user.RemainNum < 0 {
			logic.WriteRateLimitHeader(c, strconv.Itoa(0), user.ExpireTime.String())
			// return 429 to user
			c.JSON(http.StatusTooManyRequests, gin.H{
				"msg": "too many request",
			})
			c.Abort()
			return
		}
		// if not > rate limit
		b, _ := json.Marshal(&(user.RateLimitValue))
		dao.DB.Set(user.IP, string(b), user.ExpireTime.Sub(time.Now()))
		logic.WriteRateLimitHeader(c, strconv.Itoa(user.RemainNum), user.ExpireTime.String())
		// next()
		c.Next()
		return
	}
	// not has or expire: create new record to DB
	nUser := logic.CreateNewUserRateLimit()
	b, _ := json.Marshal(&(nUser.RateLimitValue))
	dao.DB.Set(user.IP, string(b), conf.RateLimitDuration)
	logic.WriteRateLimitHeader(c, strconv.Itoa(nUser.RemainNum), nUser.ExpireTime.String())
	// next()
	c.Next()
}
