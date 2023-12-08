package services_test

import (
	"testing"

	"github.com/IcaroSilvaFK/whyzy_bot_discord_api/infra/database"
	"github.com/IcaroSilvaFK/whyzy_bot_discord_api/infra/repositories"
	"github.com/IcaroSilvaFK/whyzy_bot_discord_api/infra/services"
)

func TestListAllRecommendations(t *testing.T) {

	db := database.NewDatabaseConn()
	repo := repositories.NewRecommendationsAnimesRepository(db)
	svc := services.NewRecommendationsAnimesService(repo)

	recommendations, err := svc.GetAllRecommendations()

	if err != nil {
		t.Errorf("Error on list all recommendations: %v", err)
	}

	if len(*recommendations) == 0 {
		t.Errorf("Error on list all recommendations: %v", err)
	}

}
