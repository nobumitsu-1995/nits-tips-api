package controller

import (
	"net/http"
	"nits-tips-api/model"
	"nits-tips-api/usecase"
	"os"
	"strconv"
	"time"

	"github.com/labstack/echo/v4"
)

type IReactionStampController interface {
	GetReactionStampsByArticleId(c echo.Context) error
	CreateReactionStamp(c echo.Context) error
	DeleteReactionStamp(c echo.Context) error
	getSessionData(c echo.Context) (model.SessionData, error)
	setCookie(c echo.Context, sessionId string)
	getOrCreateSessionID(c echo.Context) (string, error)
}

type reactionStampController struct {
	rsu usecase.IReactionStampUsecase
	sdu usecase.ISessionDataUsecase
}

func NewReactionStampController(rsu usecase.IReactionStampUsecase, sdu usecase.ISessionDataUsecase) IReactionStampController {
	return &reactionStampController{rsu, sdu}
}

func (rsc *reactionStampController) GetReactionStampsByArticleId(c echo.Context) error {
	articleId := c.Param("articleId")
	sessionData, err := rsc.getSessionData(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
	}
	rsc.setCookie(c, sessionData.SessionId)

	reactionStampRes, err := rsc.rsu.GetReactionStampsByArticleId(articleId, sessionData.UserId)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, reactionStampRes)
}

func (rsc *reactionStampController) CreateReactionStamp(c echo.Context) error {
	sessionData, err := rsc.getSessionData(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
	}
	rsc.setCookie(c, sessionData.SessionId)

	reactionStamp := model.ReactionStamp{}
	if err := c.Bind(&reactionStamp); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	reactionStamp.UserId = sessionData.UserId

	reactionStampResponse, err := rsc.rsu.CreateReactionStamp(reactionStamp)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusCreated, reactionStampResponse)
}

func (rsc *reactionStampController) DeleteReactionStamp(c echo.Context) error {
	reactionStampId := c.Param("reactionStampId")
	reactionStampIdUint, _ := strconv.ParseUint(reactionStampId, 10, 0)
	sessionData, err := rsc.getSessionData(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
	}
	rsc.setCookie(c, sessionData.SessionId)

	if err := rsc.rsu.DeleteReactionStamp(uint(reactionStampIdUint), sessionData.UserId); err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.NoContent(http.StatusNoContent)
}

func (rsc *reactionStampController) getSessionData(c echo.Context) (model.SessionData, error) {
	ctx := c.Request().Context()

	sessionId, err := rsc.getOrCreateSessionID(c)
	if err != nil {
		return model.SessionData{}, err
	}

	sessionData, err := rsc.sdu.GetSession(ctx, sessionId)
	if err != nil {
		return model.SessionData{}, err
	}
	return sessionData, nil
}

func (rsc *reactionStampController) setCookie(c echo.Context, sessionId string) {
	cookie := new(http.Cookie)
	cookie.Name = "sessionId"
	cookie.Value = sessionId
	cookie.Expires = time.Now().Add(24 * time.Hour * 365)
	cookie.Path = "/"
	cookie.Domain = os.Getenv("API_DOMAIN")
	cookie.Secure = true
	cookie.HttpOnly = true
	cookie.SameSite = http.SameSiteNoneMode
	c.SetCookie(cookie)
}

func (rsc *reactionStampController) getOrCreateSessionID(c echo.Context) (string, error) {
	cookie, err := c.Cookie("sessionId")
	if err == nil {
		return cookie.Value, nil
	}

	userId, err := rsc.sdu.GenerateUUID()
	if err != nil {
		return "", err
	}

	newCookie := &http.Cookie{
		Name:  "sessionId",
		Value: userId,
	}

	return newCookie.Value, nil
}
