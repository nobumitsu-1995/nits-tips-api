package db

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/redis/go-redis/v9"
)

func NewIMDB() *redis.Client {
	godotenv.Load() // env変数の読み込み

	// 接続アドレス
	address := fmt.Sprintf(
		"%s:%s",
		os.Getenv("REDIS_HOST"),
		os.Getenv("REDIS_PORT"),
	)

	imdb := redis.NewClient(&redis.Options{
		Addr:     address,
		Password: os.Getenv("REDIS_PASSWORD"),
		DB:       0,
		PoolSize: 1000,
	})

	fmt.Printf("Connected Redis PORT: %s \n", os.Getenv(("REDIS_PORT")))
	return imdb
}

func CloseIMDB(imdb *redis.Client) {
	redis := imdb.Conn()
	if err := redis.Close(); err != nil {
		log.Fatalln(err)
	}
}
