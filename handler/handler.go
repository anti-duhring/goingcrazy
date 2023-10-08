package handler

import (
	"github.com/anti-duhring/goingcrazy/config"
	"gorm.io/gorm"
)

var (
	logger *config.Logger
	db     *gorm.DB
	cache  *config.Cache
	worker *config.Worker
)

func InitializeHandler() {
	logger = config.GetLogger("handler")
	db = config.GetDB()
	cache = config.GetCache()
	worker = config.GetWorker()
}
