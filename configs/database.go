package configs

import (
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitDB() *gorm.DB {
	var (
		DB_HOST     = os.Getenv("DB_HOST")
		DB_USER     = os.Getenv("DB_USER")
		DB_PASSWORD = os.Getenv("DB_PASSWORD")
		DB_PORT     = os.Getenv("DB_PORT")
		DB_NAME     = os.Getenv("DB_NAME")
		DB_TIMEZONE = os.Getenv("DB_TIMEZONE")
		DB_SSLMODE  = os.Getenv("DB_SSLMODE")
	)

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=%s", DB_HOST, DB_USER, DB_PASSWORD, DB_NAME, DB_PORT, DB_SSLMODE, DB_TIMEZONE)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("Can't connect to database.")
	}

	return db
}
