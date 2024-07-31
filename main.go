package main

import (
	"nits-tips-api/controller"
	"nits-tips-api/db"
	"nits-tips-api/repository"
	"nits-tips-api/router"
	"nits-tips-api/usecase"
	"nits-tips-api/validator"
)

func main() {
	postgres := db.NewDB()
	redis := db.NewIMDB()

	reactionStampValidator := validator.NewReactionStampValidator()

	reactionStampRepository := repository.NewReactionStampRepository(postgres)
	sessionDataRepository := repository.NewSessionDataRepository(redis)

	reactionStampUsecase := usecase.NewReactionStampUsecase(reactionStampRepository, reactionStampValidator)
	sessionDataUsecase := usecase.NewSessionDataUsecase(sessionDataRepository)

	reactionStampController := controller.NewReactionStampController(reactionStampUsecase, sessionDataUsecase)

	e := router.NewRouter(reactionStampController)
	e.Logger.Fatal(e.Start(":80"))
}
