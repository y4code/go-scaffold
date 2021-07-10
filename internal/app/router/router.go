package router

import (
	"github.com/gin-gonic/gin"
	"go-scaffold/internal/pkg/client"
	"net/http"
)

func SetupRouter(clt *client.Client) *gin.Engine {
	r := gin.Default()

	r.GET("/examples", GetAllExamples(clt))
	r.GET("/markets", GetAllMarkets(clt))
	return r
}

func GetAllExamples(clt *client.Client) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, clt.Config.Profile)
	}
}

func GetAllMarkets(clt *client.Client) gin.HandlerFunc {
	return func(c *gin.Context) {

		if customAuth := c.GetHeader("custom_auth"); customAuth == "" {
			c.JSON(http.StatusBadRequest, gin.H{
				"code":    "40000",
				"message": "parse params error",
				"content": "",
			})
			return
		}

		allExamples := clt.ExampleService.GetAllExamples()
		c.JSON(http.StatusOK, allExamples)

	}
}
