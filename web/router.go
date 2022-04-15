package web

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func loadRouter() *gin.Engine {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.POST("/submit", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"result": "result",
		})
	})
	return r
}
