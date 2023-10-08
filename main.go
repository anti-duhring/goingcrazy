package main

import (
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/anti-duhring/goingcrazy/config"
	"github.com/anti-duhring/goingcrazy/router"
)

var (
	logger     config.Logger
	numWorkers int = 1
	numBatch   int = 1000
)

func main() {
	logger = *config.GetLogger("main")

	time.Sleep(5 * time.Second) // Database on compose
	err := config.Init()
	if err != nil {
		logger.Errorf("Config initialization error: %v", err)

		return
	}

	chExit := make(chan struct{})

	for i := 0; i < numWorkers; i++ {
		go config.RunWorker(config.GetWorker().ChPeople, chExit, numBatch)
	}

	router.Initialize()

	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM)
	for i := 0; i < numWorkers; i++ {
		<-chExit
	}
}
