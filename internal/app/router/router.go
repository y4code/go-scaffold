package router

import (
	"github.com/gin-gonic/gin"
	"github.com/swaggo/swag"
	"go-scaffold/internal/pkg/client"
	"net/http"
)

var clt *client.Client

func SetupRouter(client *client.Client) *gin.Engine {
	clt = client
	r := gin.Default()
	r.GET("/swagger.json", func(c *gin.Context) {
		swaggerfile, _ := swag.ReadDoc()
		c.String(200, "%s", swaggerfile)
	})
	r.GET("/example", GetExample)
	return r
}

func GetExample(c *gin.Context) {
	c.JSON(http.StatusOK, clt.Config.Profile)
}
