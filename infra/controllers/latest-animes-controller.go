package controllers

import (
	"net/http"

	"github.com/IcaroSilvaFK/whyzy_bot_discord_api/infra/services"
	"github.com/labstack/echo/v4"
)

type LatestAnimesController struct {
	svc services.LatestAnimesServiceInterface
}

type LatestAnimesControllerInterface interface {
	GetLatestAnimes(ctx echo.Context) error
}

func NewLatestAnimesController(
	svc services.LatestAnimesServiceInterface,
) LatestAnimesControllerInterface {
	return &LatestAnimesController{
		svc,
	}
}

func (lc *LatestAnimesController) GetLatestAnimes(ctx echo.Context) error {

	animes, err := lc.svc.GetLatestAnimes()

	if err != nil {
		return err
	}

	return ctx.JSON(http.StatusOK, animes)

}
