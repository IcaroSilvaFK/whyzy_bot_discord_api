package di

import (
	"github.com/IcaroSilvaFK/whyzy_bot_discord_api/infra/controllers"
	"github.com/IcaroSilvaFK/whyzy_bot_discord_api/infra/database"
	"github.com/IcaroSilvaFK/whyzy_bot_discord_api/infra/repositories"
	"github.com/IcaroSilvaFK/whyzy_bot_discord_api/infra/services"
)

func NewLatestAnimesControllerDI() controllers.LatestAnimesControllerInterface {

	db := database.NewDatabaseConn()
	repo := repositories.NewLatestAnimesRepository(db)
	svc := services.NewLatestAnimesService(repo)

	return controllers.NewLatestAnimesController(svc)
}
