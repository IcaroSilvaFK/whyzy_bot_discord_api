package repositories_test

import (
	"database/sql"
	"testing"

	"github.com/IcaroSilvaFK/whyzy_bot_discord_api/infra/database"
	"github.com/IcaroSilvaFK/whyzy_bot_discord_api/infra/repositories"
)

func TestCreateRecommendations(t *testing.T) {

	db := database.NewDatabaseConn()
	lsr := repositories.NewRecommendationsAnimesRepository(db)

	recommendations, err := lsr.CreateRecommendations()

	if err != nil {
		t.Errorf("Error on create recommendations: %v", err)
	}

	if len(*recommendations) == 0 {
		t.Errorf("Error on create recommendations: %v", err)
	}

}

func TestListAllRecommendations(t *testing.T) {

	db := database.NewDatabaseConn()
	lsr := repositories.NewRecommendationsAnimesRepository(db)

	_, err := lsr.GetAllRecommendations()

	if err != nil {
		t.Errorf("Error on list all recommendations: %v", err)
	}

	if err == sql.ErrNoRows {
		lsr.CreateRecommendations()
	}

	recommendations, err := lsr.GetAllRecommendations()

	if err != nil {
		t.Errorf("Error on list all recommendations: %v", err)
	}

	if len(*recommendations) == 0 {
		t.Errorf("Error on list all recommendations: %v", err)
	}

}
