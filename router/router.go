package router

import (
	"github.com/gin-gonic/gin"
	"github.com/jamadeu/api-account/handler"
	"github.com/jamadeu/api-account/repository"
	"gorm.io/gorm"
)

type Router struct {
	port string
	db   *gorm.DB
}

func NewRouter(addr string, db *gorm.DB) *Router {
	return &Router{
		port: addr,
		db:   db,
	}
}

func (r *Router) Initialize() error {
	// Initialize router
	router := gin.Default()

	ar := repository.NewAccountRepository(r.db)
	ah := handler.NewAccountHandler(ar)
	// Initialize routes
	InitializeRoutes(router, ah)

	// Run the server
	return router.Run(r.port)
}
