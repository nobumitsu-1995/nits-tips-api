package db

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewDB() *gorm.DB {
	godotenv.Load() // env変数の読み込み

	// DB接続URL
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s",
		os.Getenv(("POSTGRES_HOST")),
		os.Getenv(("POSTGRES_USER")),
		os.Getenv(("POSTGRES_PASSWORD")),
		os.Getenv(("POSTGRES_DB")),
		os.Getenv(("POSTGRES_PORT")),
	)

	// DBと接続
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Printf("Connected Postgres PORT: %s", os.Getenv(("POSTGRES_PORT")))
	return db
}

func CloseDB(db *gorm.DB) {
	postgres, _ := db.DB()
	if err := postgres.Close(); err != nil {
		log.Fatalln(err)
	}
}
