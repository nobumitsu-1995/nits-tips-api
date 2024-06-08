package main

import (
	"fmt"
	"nits-tips-api/db"
	"nits-tips-api/model"
)

func main() {
	postgres := db.NewDB()

	defer fmt.Print("migrationが正常に実行されました")
	defer db.CloseDB(postgres)

	postgres.AutoMigrate(&model.ReactionStamp{})
}
