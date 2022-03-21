package main

import (
	"net/http"

	"github.com/Sxu-Online-Judge/judger/judge"
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
		submit := judge.Submit{}
		if err := c.ShouldBindJSON(&submit); err != nil {
			c.JSON(http.StatusOK, gin.H{
				"msg": "bind model error",
			})
			return
		}

		compiler := judge.NewCompiler(&submit)
		compileResult := compiler.Run()
		if compileResult.Status == judge.StatusCE {
			c.JSON(http.StatusOK, gin.H{
				"result": compileResult,
			})
		}

		judger := judge.NewJudger(&submit)
		// judger.Print()
		result := judger.Judge()

		c.JSON(http.StatusOK, gin.H{
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
