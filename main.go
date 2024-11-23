package main

import (
	"github.com/jamadeu/api-account/config"
	"github.com/jamadeu/api-account/router"
)

func main() {
	logger := config.GetLogger()

	// Initialize configs
	if err := config.Init(); err != nil {
		logger.Error(err.Error())
		return
	}
	r := router.NewRouter(":8080", config.GetConnDB())
	if err := r.Initialize(); err != nil {
		panic(err)
	}
}
