package controller

import (
	"net/http"
	"nits-tips-api/model"
	"nits-tips-api/usecase"
	"strconv"

	"github.com/labstack/echo/v4"
)

type IReactionStampController interface {
	GetReactionStampsByArticleId(c echo.Context) error
	CreateReactionStamp(c echo.Context) error
	DeleteReactionStamp(c echo.Context) error
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

func (rsc *reactionStampController) CreateReactionStamp(c echo.Context) error {

	reactionStamp := model.ReactionStamp{}
	if err := c.Bind(&reactionStamp); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	reactionStampResponse, err := rsc.rsu.CreateReactionStamp(reactionStamp)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusCreated, reactionStampResponse)
}

func (rsc *reactionStampController) DeleteReactionStamp(c echo.Context) error {
	reactionStampId := c.Param("reactionStampId")
	userId := c.Param("userId")
	reactionStampIdUint, _ := strconv.ParseUint(reactionStampId, 10, 0)
	userIdUint, _ := strconv.ParseUint(userId, 10, 0)

	err := rsc.rsu.DeleteReactionStamp(uint(reactionStampIdUint), uint(userIdUint))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.NoContent(http.StatusNoContent)
}
