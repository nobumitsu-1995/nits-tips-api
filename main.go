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
	redis := db.NewIMDB()

	reactionStampRepository := repository.NewReactionStampRepository(postgres)
	sessionDataRepository := repository.NewSessionDataRepository(redis)

	reactionStampUsecase := usecase.NewReactionStampUsecase(reactionStampRepository)
	sessionDataUsecase := usecase.NewSessionDataUsecase(sessionDataRepository)

	reactionStampController := controller.NewReactionStampController(reactionStampUsecase, sessionDataUsecase)

	e := router.NewRouter(reactionStampController)
	e.Logger.Fatal(e.Start(":80"))
}
