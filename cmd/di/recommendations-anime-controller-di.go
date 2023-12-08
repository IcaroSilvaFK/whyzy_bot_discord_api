package di

import (
	"github.com/IcaroSilvaFK/whyzy_bot_discord_api/infra/controllers"
	"github.com/IcaroSilvaFK/whyzy_bot_discord_api/infra/database"
	"github.com/IcaroSilvaFK/whyzy_bot_discord_api/infra/repositories"
	"github.com/IcaroSilvaFK/whyzy_bot_discord_api/infra/services"
)

func NewRecommendationsAnimesControllerDI() controllers.RecommendationsControllerInterface {

	db := database.NewDatabaseConn()
	repo := repositories.NewRecommendationsAnimesRepository(db)
	svc := services.NewRecommendationsAnimesService(repo)

	return controllers.NewRecommendationsController(svc)
}
