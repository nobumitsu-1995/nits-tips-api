package controller

import (
	"net/http"
	"nits-tips-api/usecase"

	"github.com/labstack/echo/v4"
)

type IReactionStampController interface {
	GetReactionStampsByArticleId(c echo.Context) error
}

type reactionStampController struct {
	rsu usecase.IReactionStampUsecase
}

func NewReactionStampController(rsu usecase.IReactionStampUsecase) IReactionStampController {
	return &reactionStampController{rsu}
}

func (rsc *reactionStampController) GetReactionStampsByArticleId(c echo.Context) error {
	articleId := c.Param("articleId")

	reactionStampRes, err := rsc.rsu.GetReactionStampsByArticleId(articleId)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, reactionStampRes)
}
