package services_test

import (
	"testing"

	"github.com/IcaroSilvaFK/whyzy_bot_discord_api/infra/database"
	"github.com/IcaroSilvaFK/whyzy_bot_discord_api/infra/repositories"
	"github.com/IcaroSilvaFK/whyzy_bot_discord_api/infra/services"
)

func TestGetLatestAnimes(t *testing.T) {
	db := database.NewDatabaseConn()
	repo := repositories.NewLatestAnimesRepository(db)
	svc := services.NewLatestAnimesService(repo)

	animes, err := svc.GetLatestAnimes()

	if err != nil {
		t.Errorf("Error on get latest animes: %v", err)
	}

	if len(*animes) == 0 {
		t.Errorf("Error on get latest animes: %v", err)
	}

}
