package main

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

// LimitDuration : 造訪上限時間(單位秒)
const LimitDuration = 60

// RateLimit : 限定時間的造訪上限
const RateLimit = 3

func main() {
	r := gin.Default()
	r.LoadHTMLGlob("views/*")

	r.GET("/", rateLimetMid, homeHandler)
	r.GET("/hello", rateLimetMid, helloHandler)

	r.Run() // port :8086
}

func homeHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", nil)
}

func helloHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "hello.html", nil)
}

func rateLimetMid(c *gin.Context) {
	// 取出cookie中timer資訊
	timer, err := c.Cookie("rate-limit-timer")
	// 如果第一次造訪(or過期)會沒有cookie則 : 設定timer就設定timer以及counter
	if err != nil {
		expire := time.Now().Add(time.Second * LimitDuration).String() // 過期時間為now + 1 hour
		c.SetCookie("rate-limit-timer", expire, LimitDuration, "/", "127.0.0.1", false, true)
		c.SetCookie("rate-limit-counter", strconv.Itoa(RateLimit-1), LimitDuration, "/", "127.0.0.1", false, true)
		c.Writer.Header().Set("X-RateLimit-Remaining", strconv.Itoa(RateLimit-1))
		c.Writer.Header().Set("X-RateLimit-Reset", expire)
		// 往下一個func
		c.Next()
		return
	}

	// 取出cookie中 剩餘次數 資訊
	ck, _ := c.Cookie("rate-limit-counter")
	remainTimes, _ := strconv.Atoi(ck)
	// 如果已達限制數量
	if remainTimes < 1 {
		c.Writer.Header().Set("X-RateLimit-Remaining", strconv.Itoa(remainTimes))
		c.Writer.Header().Set("X-RateLimit-Reset", timer)
		c.JSON(http.StatusTooManyRequests, nil)
		c.Abort()
		return
	}

	// 未達限制數量: 設定cookie中 > 新的剩餘次數
	newRemain := strconv.Itoa(remainTimes - 1)
	c.SetCookie("rate-limit-counter", newRemain, LimitDuration, "/", "127.0.0.1", false, true)
	// X-RateLimit-Remaining: 剩餘的請求數量
	c.Writer.Header().Set("X-RateLimit-Remaining", newRemain)
	// X-RateLimit-Reset : rate limit 歸零的時間
	c.Writer.Header().Set("X-RateLimit-Reset", timer)
	// 往下一個func
	c.Next()
}
