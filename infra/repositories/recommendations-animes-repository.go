package repositories

import (
	"database/sql"
	"encoding/json"
	"io"
	"log"
	"net/http"

	"github.com/IcaroSilvaFK/whyzy_bot_discord_api/infra/entities"
	entities_external "github.com/IcaroSilvaFK/whyzy_bot_discord_api/infra/entities/jikan"
)

type RecommendationsAnimesRepository struct {
	db *sql.DB
}

type RecommendationsAnimesRepositoryInterface interface {
	GetAllRecommendations() (*[]entities.RecommendationsAnimesEntity, error)
	CreateRecommendations() (*[]entities.RecommendationsAnimesEntity, error)
}

func NewRecommendationsAnimesRepository(db *sql.DB) RecommendationsAnimesRepositoryInterface {
	return &RecommendationsAnimesRepository{db}
}

func (repo *RecommendationsAnimesRepository) GetAllRecommendations() (*[]entities.RecommendationsAnimesEntity, error) {

	result, err := repo.db.Query("SELECT * FROM recommendations")

	if err != nil {
		return nil, err
	}

	defer result.Close()

	var recommendations []entities.RecommendationsAnimesEntity

	for result.Next() {

		var recommendation entities.RecommendationsAnimesEntity

		err := result.Scan(&recommendation.MalId, &recommendation.Content, &recommendation.Title, &recommendation.ImageUrl, &recommendation.Url)

		if err != nil {
			return nil, err
		}

		recommendations = append(recommendations, recommendation)

	}

	return &recommendations, nil
}

func (rep *RecommendationsAnimesRepository) CreateRecommendations() (*[]entities.RecommendationsAnimesEntity, error) {

	resp, err := http.Get("https://api.jikan.moe/v4/recommendations/anime?limit=1")

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	bt, err := io.ReadAll(resp.Body)

	if err != nil {
		return nil, err
	}

	var jikanEntity entities_external.RecommendationsAnimesDataJikanApi

	err = json.Unmarshal(bt, &jikanEntity)

	if err != nil {
		return nil, err
	}

	var recommendations []entities.RecommendationsAnimesEntity

	idx := 0

recommendations:
	for _, v := range jikanEntity.Data {
		for _, recommendation := range v.Entry {

			if idx > 5 {
				break recommendations
			}

			recommendations = append(recommendations, entities.RecommendationsAnimesEntity{
				Title:    recommendation.Title,
				MalId:    recommendation.MalId,
				Content:  v.Content,
				Url:      recommendation.Url,
				ImageUrl: recommendation.Image.JPG.ImageUrl,
			})
			idx++
		}
	}

	go func() {
		for _, recommendation := range recommendations {
			_, err := rep.db.Exec("INSERT INTO recommendations VALUES (?, ?, ?, ?, ?)", recommendation.MalId, recommendation.Content, recommendation.Title, recommendation.ImageUrl, recommendation.Url)

			log.Println(err)
		}
	}()

	return &recommendations, nil
}
