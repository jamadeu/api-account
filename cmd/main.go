package main

import (
	"github.com/jamadeu/api-account/config"
	"github.com/jamadeu/api-account/router"
)

func main() {
	logger := config.GetLogger()

	// Initialize configs
	if err := config.Init(); err != nil {
		logger.Panic(err.Error())
		return
	}
	router.Initialize()
}
