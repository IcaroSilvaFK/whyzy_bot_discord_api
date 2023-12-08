package services

import (
	"database/sql"

	"github.com/IcaroSilvaFK/whyzy_bot_discord_api/infra/entities"
	"github.com/IcaroSilvaFK/whyzy_bot_discord_api/infra/repositories"
)

type LatestAnimesService struct {
	repo repositories.LatestAnimesRepositoryInterface
}

type LatestAnimesServiceInterface interface {
	GetLatestAnimes() (*[]entities.LatestAnimesEntity, error)
}

func NewLatestAnimesService(
	repo repositories.LatestAnimesRepositoryInterface,
) LatestAnimesServiceInterface {

	return &LatestAnimesService{
		repo,
	}
}

func (svc *LatestAnimesService) GetLatestAnimes() (*[]entities.LatestAnimesEntity, error) {

	animes, err := svc.repo.GetLatestAnimesAvailable()

	if err != nil && err != sql.ErrNoRows {
		return nil, err
	}

	if err == sql.ErrNoRows || len(*animes) == 0 {
		animes, err := svc.repo.CreateLatestAnimes()

		if err != nil {
			return nil, err
		}

		return animes, nil
	}

	return animes, nil
}
