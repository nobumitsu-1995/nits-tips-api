package main

import (
	"nits-tips-api/controller"
	"nits-tips-api/db"
	"nits-tips-api/repository"
	"nits-tips-api/router"
	"nits-tips-api/usecase"
)

func main() {
	postgres := db.NewDB()
	db.NewIMDB()

	reactionStampRepository := repository.NewReactionStampRepository(postgres)
	reactionStampUsecase := usecase.NewReactionStampUsecase(reactionStampRepository)
	reactionStampController := controller.NewReactionStampController(reactionStampUsecase)

	e := router.NewRouter(reactionStampController)
	e.Logger.Fatal(e.Start(":80"))
}
