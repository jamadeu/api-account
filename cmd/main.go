package main

import (
	"github.com/gin-gonic/gin"
	"github.com/jamadeu/api-account/config"
)

func main() {
	logger := config.GetLogger()

	// Initialize configs
	if err := config.Init(); err != nil {
		logger.Panic(err.Error())
		return
	}
	s := gin.Default()
	s.Run(":8000")
}
