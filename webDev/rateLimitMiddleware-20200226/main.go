package main

import (
	"rateLimitMiddleware/conf"
	"rateLimitMiddleware/controller"
	"rateLimitMiddleware/dao"

	"github.com/gin-gonic/gin"
)

func main() {
	// connect to Database
	err := dao.ConnectDataBase()
	if err != nil {
		panic(err)
	}
	defer dao.DB.Close()

	// init router
	r := gin.Default()
	// set index page api, use the middleware
	r.GET("/", controller.RateLimitMiddleware, controller.HomeHandler)

	// start service
	r.Run(":" + conf.ServicePort)
}
