package router

import (
	"nits-tips-api/controller"

	"github.com/labstack/echo/v4"
)

func NewRouter(rsc controller.IReactionStampController) *echo.Echo {
	e := echo.New()

	rs := e.Group("/reactionStamps")
	rs.GET("/:articleId", rsc.GetReactionStampsByArticleId)
	rs.POST("", rsc.CreateReactionStamp)
	rs.DELETE("/:reactionStampId/:userId", rsc.DeleteReactionStamp)

	return e
}
