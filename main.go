package main

import (
	"github.com/anti-duhring/goingcrazy/config"
	"github.com/anti-duhring/goingcrazy/router"
)

var (
	logger config.Logger
)

func main() {
	logger = *config.GetLogger("main")

	err := config.Init()
	if err != nil {
		logger.Errorf("Config initialization error: %v", err)

		return
	}

	router.Initialize()
}
