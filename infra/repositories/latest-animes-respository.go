package repositories

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/IcaroSilvaFK/whyzy_bot_discord_api/infra/entities"
	entities_external "github.com/IcaroSilvaFK/whyzy_bot_discord_api/infra/entities/jikan"
)

type LatestAnimesRespository struct {
	db *sql.DB
}

type LatestAnimesRepositoryInterface interface {
	GetLatestAnimesAvailable() (*[]entities.LatestAnimesEntity, error)
	CreateLatestAnimes() (*[]entities.LatestAnimesEntity, error)
}

func NewLatestAnimesRepository(
	db *sql.DB,
) LatestAnimesRepositoryInterface {
	return &LatestAnimesRespository{
		db,
	}
}

func (lsr *LatestAnimesRespository) GetLatestAnimesAvailable() (*[]entities.LatestAnimesEntity, error) {

	r, err := lsr.db.Query("SELECT * FROM latest_animes")

	if err != nil {
		return nil, err
	}

	defer r.Close()

	var latestAnimes []entities.LatestAnimesEntity

	for r.Next() {

		var entity entities.LatestAnimesEntity

		err := r.Scan(
			&entity.MalId,
			&entity.Url,
			&entity.Title,
			&entity.Rank,
			&entity.ImageUrl,
		)

		if err != nil {
			return nil, err
		}
		latestAnimes = append(latestAnimes, entity)
	}

	if len(latestAnimes) == 0 {

		return nil, sql.ErrNoRows
	}

	return &latestAnimes, nil
}

func (lsr *LatestAnimesRespository) CreateLatestAnimes() (*[]entities.LatestAnimesEntity, error) {

	resp, err := http.Get("https://api.jikan.moe/v4/top/anime?limit=10")

	if err != nil {

		return nil, err
	}

	defer resp.Body.Close()

	bt, err := io.ReadAll(resp.Body)

	if err != nil {

		return nil, err
	}

	var latestAnimes entities_external.LatestAnimesEntityJikanApi

	err = json.Unmarshal(bt, &latestAnimes)

	fmt.Println(err)

	if err != nil {

		return nil, err
	}

	for _, anime := range latestAnimes.Data {
		_, err = lsr.db.Exec(
			"INSERT INTO latest_animes (mal_id, url, title, rank, image_url) VALUES (?, ?, ?, ?, ?)",
			anime.MalId,
			anime.Url,
			anime.Title,
			anime.Rank,
			anime.Images.JPG.ImageUrl,
		)

		if err != nil {

			return nil, err
		}
	}

	var animes []entities.LatestAnimesEntity

	for _, anime := range latestAnimes.Data {

		animes = append(animes, entities.LatestAnimesEntity{
			MalId:    anime.MalId,
			Url:      anime.Url,
			Title:    anime.Title,
			Rank:     anime.Rank,
			ImageUrl: anime.Images.JPG.ImageUrl,
		})

	}

	return &animes, nil
}
