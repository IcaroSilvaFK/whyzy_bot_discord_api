package controllers

import (
	"net/http"

	"github.com/IcaroSilvaFK/whyzy_bot_discord_api/infra/services"
	"github.com/labstack/echo/v4"
)

type RecommendationsController struct {
	svc services.RecommendationsAnimesServiceInterface
}

type RecommendationsControllerInterface interface {
	GetRecommendations(ctx echo.Context) error
}

func NewRecommendationsController(
	svc services.RecommendationsAnimesServiceInterface,
) RecommendationsControllerInterface {
	return &RecommendationsController{
		svc,
	}
}

func (rc *RecommendationsController) GetRecommendations(ctx echo.Context) error {

	animes, err := rc.svc.GetAllRecommendations()

	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, echo.Map{
			"error": err.Error(),
		})
	}

	return ctx.JSON(http.StatusOK, animes)
}
