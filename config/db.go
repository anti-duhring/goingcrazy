package config

import (
	"os"

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

	// dbError = db.AutoMigrate(&schema.Person{})
	// logger.Info("Automigrating database...")

	// db.Exec(`
	// 	CREATE EXTENSION IF NOT EXISTS pg_trgm;
	// 	CREATE INDEX IF NOT EXISTS idx_people_search ON public.people USING gist (search_index public.gist_trgm_ops (siglen='64'));
	// `)

	sqlDB, err := db.DB()
	if err != nil {
		logger.Errorf("Database connection error: %v", err)
		return nil, err
	}

	sqlDB.SetMaxOpenConns(15)

	return db, dbError
}
