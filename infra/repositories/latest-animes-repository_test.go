package repositories_test

import (
	"database/sql"
	"testing"

	"github.com/IcaroSilvaFK/whyzy_bot_discord_api/infra/database"
	"github.com/IcaroSilvaFK/whyzy_bot_discord_api/infra/repositories"
)

func TestCreateLatestAnimesAVailable(t *testing.T) {
	db := database.NewDatabaseConn()
	lsr := repositories.NewLatestAnimesRepository(db)

	animes, err := lsr.CreateLatestAnimes()

	if err != nil {
		t.Errorf("Error on create latest animes available: %v", err)
	}

	if len(*animes) == 0 {
		t.Errorf("Error on create latest animes available: %v", err)
	}
}

func TestGetLatestAnimesAvailable(t *testing.T) {
	db := database.NewDatabaseConn()
	lsr := repositories.NewLatestAnimesRepository(db)

	_, err := lsr.GetLatestAnimesAvailable()

	if err != nil && err != sql.ErrNoRows {
		t.Errorf("Error on get latest animes available: %v", err)
	}

	if err == sql.ErrNoRows {
		_, err := lsr.CreateLatestAnimes()

		if err != nil {
			t.Errorf("Error on get latest animes available: %v", err)
		}
	}
	animes, err := lsr.GetLatestAnimesAvailable()

	if err != nil {
		t.Errorf("Error on get latest animes available: %v", err)
	}

	if len(*animes) == 0 {
		t.Errorf("Error on get latest animes available: %v", err)
	}

}
