package config

import (
	"github.com/joho/godotenv"
	"gorm.io/gorm"
	"os"
)

type CustomError struct {
	Message string
}

func (ce CustomError) Error() string {
	return ce.Message
}

func GetEnv(key string) string {
	err := godotenv.Load(".env")
	if err != nil {
		panic("error loading .env file !")
	}
	return os.Getenv(key)
}

// CountWithoutPreload is Scope to calculate data without Preload
func CountWithoutPreload(db *gorm.DB) *gorm.DB {
	return db.Select("COUNT(*)")
}
