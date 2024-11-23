package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jamadeu/api-account/docs"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func initializeRoutes(router *gin.Engine) {

	basePath := "/api/v1"
	docs.SwaggerInfo.BasePath = basePath
	v1 := router.Group(basePath)
	{
		v1.GET("", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"message": "test ok"})
		})
	}

	// Initialize Swagger
	router.GET("/swagger/*any",
		ginSwagger.WrapHandler(swaggerfiles.Handler))
}
