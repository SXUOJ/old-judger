package main

import (
	"net/http"

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
		test := model.Submit{}
		if err := c.ShouldBindJSON(&test); err != nil {
			c.JSON(http.StatusOK, gin.H{
				"msg": "bind model error",
			})
			return
		}

		//TODO:judge

		c.JSON(http.StatusOK, gin.H{
			"submit_id":        test.SubmitId,
			"problem_id":       test.ProblemId,
			"problem_type":     test.ProblemType,
			"code_type":        test.CodeType,
			"code_source_path": test.CodeSourcePath,
			"time_limit":       test.TimeLimit,
			"memory_limit":     test.MemoryLimit,
		})
	})
	return r
}

func main() {
	r := setupRouter()
	r.Run()
}
