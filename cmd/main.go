package main

import (
	"net/http"

	router "github.com/IcaroSilvaFK/whyzy_bot_discord_api/cmd/routes"
	"github.com/IcaroSilvaFK/whyzy_bot_discord_api/infra/database"
	"github.com/IcaroSilvaFK/whyzy_bot_discord_api/infra/repositories"
	"github.com/labstack/echo/v4"
)

func main() {

	db := database.NewDatabaseConn()

	svc := repositories.NewLatestAnimesRepository(db)

	defer db.Close()

	svc.CreateLatestAnimes()

	e := echo.New()

	group := e.Group("/whyzy/api/v1")

	group.GET("/heath", func(ctx echo.Context) error {
		return ctx.JSON(http.StatusOK, echo.Map{
			"message": "ok",
			"status":  http.StatusOK,
			"method":  ctx.Request().Method,
		})
	})
	router.NewApplicationRoutes(group)

	e.Logger.Error(e.Start(":8080"))
}
