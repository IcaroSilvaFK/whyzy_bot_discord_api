package services

import (
	"database/sql"
	"fmt"

	"github.com/IcaroSilvaFK/whyzy_bot_discord_api/infra/entities"
	"github.com/IcaroSilvaFK/whyzy_bot_discord_api/infra/repositories"
)

type RecommendationsAnimesService struct {
	repo repositories.RecommendationsAnimesRepositoryInterface
}

type RecommendationsAnimesServiceInterface interface {
	GetAllRecommendations() (*[]entities.RecommendationsAnimesEntity, error)
}

func NewRecommendationsAnimesService(repo repositories.RecommendationsAnimesRepositoryInterface) RecommendationsAnimesServiceInterface {

	return &RecommendationsAnimesService{
		repo: repo,
	}
}

func (svc *RecommendationsAnimesService) GetAllRecommendations() (*[]entities.RecommendationsAnimesEntity, error) {

	animes, err := svc.repo.GetAllRecommendations()

	fmt.Println(animes, err)

	if err != nil && err != sql.ErrNoRows {
		return nil, err
	}

	if err == sql.ErrNoRows || len(*animes) == 0 {
		animes, err = svc.repo.CreateRecommendations()

		if err != nil {
			return nil, err
		}

		return animes, nil
	}

	return animes, nil
}
