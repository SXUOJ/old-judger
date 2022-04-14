package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/isther/judger/judge"
	"github.com/isther/judger/model"
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
		compileResult := judger.Compiler.Run()
		if compileResult.Status != strconv.FormatInt(judge.SUCCEED, 10) {
			compileResult.Status = judge.GetJudgeStatus(strconv.FormatInt(judge.StatusCE, 10))
			c.JSON(http.StatusOK, gin.H{
				"result": compileResult,
			})
			return
		}

		runResult := judger.Runner.Run()

		c.JSON(http.StatusOK, gin.H{
			"result": runResult,
		})
	})
	return r
}

//TODO: 1. add log
func main() {
	r := setupRouter()
	r.Run(":9000")
}
