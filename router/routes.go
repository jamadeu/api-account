package router

import (
	"github.com/gin-gonic/gin"
	"github.com/jamadeu/api-account/docs"
	"github.com/jamadeu/api-account/handler"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func InitializeRoutes(router *gin.Engine, ah *handler.AccountHandler) {
	basePath := "/api/v1"
	docs.SwaggerInfo.BasePath = basePath
	ah.RegisterRoutes(router, basePath)

	// Initialize Swagger
	router.GET("/swagger/*any",
		ginSwagger.WrapHandler(swaggerfiles.Handler))
}
