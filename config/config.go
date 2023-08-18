package config

import (
	"errors"

	"gorm.io/gorm"
)

var (
	db *gorm.DB
)

func Init() error {
	return errors.New("fake error")
}

func GetLogger(p string) *Logger {
	logger := NewLogger(p)

	return logger
}
