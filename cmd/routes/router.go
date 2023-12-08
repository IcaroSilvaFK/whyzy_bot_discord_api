package router

import (
	"net/http"

	"github.com/IcaroSilvaFK/whyzy_bot_discord_api/cmd/di"
	"github.com/labstack/echo/v4"
)

func NewApplicationRoutes(r *echo.Group) {

	cAnimes := di.NewLatestAnimesControllerDI()
	cRecommendations := di.NewRecommendationsAnimesControllerDI()

	r.GET("/animes/latest", func(ctx echo.Context) error {
		return ctx.JSON(http.StatusOK, echo.Map{
			"message": "ok",
			"status":  http.StatusOK,
			"method":  ctx.Request().Method,
		})
	})

	r.GET("/animes/latest", cAnimes.GetLatestAnimes)
	r.GET("/animes/recommendations", cRecommendations.GetRecommendations)

}
