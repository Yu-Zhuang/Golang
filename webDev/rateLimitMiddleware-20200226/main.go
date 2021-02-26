package main

import (
	"rateLimitMiddleware/conf"
	"rateLimitMiddleware/controller"
	"rateLimitMiddleware/dao"

	"github.com/gin-gonic/gin"
)

func main() {
	// connect to Database
	dao.ConnectDataBase()
	defer dao.DB.Close()

	// init router
	r := gin.Default()

	r.GET("/", controller.RateLimitMiddleware, controller.HomeHandler)

	r.Run(":" + conf.ServicePort)
}
