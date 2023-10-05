package config

import (
	"os"

	"github.com/anti-duhring/goingcrazy/schema"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitializeDB() (*gorm.DB, error) {

	logger := GetLogger("database")

	envError := godotenv.Load()

	if envError != nil {
		logger.Errorf("Env loading error: %v", envError)
	}

	dsn := os.Getenv("DATABASE_URL")
	db, dbError := gorm.Open(postgres.Open(dsn), &gorm.Config{TranslateError: true})

	if dbError != nil {
		logger.Errorf("Database opening error: %v", dbError)

		return nil, dbError
	}

	dbError = db.AutoMigrate(&schema.Person{})
	logger.Info("Automigrating database...")

	if dbError != nil {
		logger.Errorf("Database automigration error: %v", dbError)
		return nil, dbError
	}

	return db, dbError
}
