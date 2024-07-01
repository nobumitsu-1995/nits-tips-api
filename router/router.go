package router

import (
	"nits-tips-api/controller"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func NewRouter(rsc controller.IReactionStampController) *echo.Echo {
	e := echo.New()

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins:     []string{"http://localhost:4321", os.Getenv("FE_URL")},
		AllowHeaders:     []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept, echo.HeaderAccessControlAllowHeaders, echo.HeaderXCSRFToken},
		AllowMethods:     []string{"GET", "PUT", "POST", "DELETE"},
		AllowCredentials: true,
	}))

	rs := e.Group("/reactionStamps")
	rs.GET("/:articleId", rsc.GetReactionStampsByArticleId)
	rs.POST("", rsc.CreateReactionStamp)
	rs.DELETE("/:reactionStampId", rsc.DeleteReactionStamp)

	return e
}
