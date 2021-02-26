package logic

import (
	"rateLimitMiddleware/conf"
	"rateLimitMiddleware/models"
	"time"

	"github.com/gin-gonic/gin"
)

// CreateNewUserRateLimit ...
func CreateNewUserRateLimit() models.RateLimit {
	var user models.RateLimit
	user.ExpireTime = time.Now().Add(conf.RateLimitDuration)
	user.RemainNum = conf.RateLimitNum - 1
	return user
}

// WriteRateLimitHeader ...
func WriteRateLimitHeader(c *gin.Context, remaining string, reset string) {
	c.Writer.Header().Set("X-RateLimit-Remaining", remaining)
	c.Writer.Header().Set("X-RateLimit-Reset", reset)
}
