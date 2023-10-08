package config

import (
	"fmt"

	"gorm.io/gorm"
)

var (
	db     *gorm.DB
	cache  *Cache
	worker *Worker
	logger *Logger
)

func Init() error {
	var err error

	db, err = InitializeDB()

	if err != nil {
		return fmt.Errorf("Error initializing DB: %v", err)
	}

	cache, err = InitializeCache()

	if err != nil {
		return fmt.Errorf("Error initializing cache: %v", err)
	}

	worker = NewWorker(db, cache.Client)

	return nil
}

func GetDB() *gorm.DB {
	return db
}

func GetCache() *Cache {
	return cache
}

func GetWorker() *Worker {
	return worker
}

func GetLogger(p string) *Logger {
	logger = NewLogger(p)

	return logger
}
