package main

import (
	"net/http"
	"sendREST/controller"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "<h1>hello world</h1>")
	})
	// get fhir data
	r.GET("/:type/:id", func(c *gin.Context) {
		t, id := c.Param("type"), c.Param("id")
		res := controller.GetFhir(t, id)
		c.String(http.StatusOK, res)
	})

	r.GET("/:type", func(c *gin.Context) {
		t := c.Param("type")
		res := controller.GetFhirAll(t)
		c.String(http.StatusOK, res)
	})
	r.Run()
}
