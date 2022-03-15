package main

import (
	"net/http"

	"github.com/Sxu-Online-Judge/judger/judge"
	"github.com/Sxu-Online-Judge/judger/model"
	"github.com/gin-gonic/gin"
)

func setupRouter() *gin.Engine {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.POST("/submit", func(c *gin.Context) {
		submit := model.Submit{}
		if err := c.ShouldBindJSON(&submit); err != nil {
			c.JSON(http.StatusOK, gin.H{
				"msg": "bind model error",
			})
			return
		}

		judger := judge.NewJudger(&submit)
		// judger.Print()
		result := judger.Judge()

		c.JSON(http.StatusOK, gin.H{
			"msg":    "good",
			"result": result,
		})
	})
	return r
}

//TODO: 1. add log
func main() {
	r := setupRouter()
	r.Run()
}
