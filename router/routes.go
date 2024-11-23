package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func initializeRoutes(router *gin.Engine) {

	basePath := "/api/v1"
	v1 := router.Group(basePath)
	{
		v1.GET("", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"message": "test ok"})
		})
	}
}
