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

	// DBと接続
	db, err := gorm.Open(postgres.Open(os.Getenv("DATABASE_URL")), &gorm.Config{})
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Printf("Connected Postgres PORT: %s \n", os.Getenv(("POSTGRES_PORT")))
	return db
}

func CloseDB(db *gorm.DB) {
	postgres, _ := db.DB()
	if err := postgres.Close(); err != nil {
		log.Fatalln(err)
	}
}
