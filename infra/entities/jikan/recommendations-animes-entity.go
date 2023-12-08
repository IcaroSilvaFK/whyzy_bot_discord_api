package entities_external

type RecommendationsAnimesDataJikanApi struct {
	Data []RecommendationsAnimesDataJikan `json:"data"`
}

type RecommendationsAnimesDataJikan struct {
	Entry   []RecommendationsAnimesData `json:"entry"`
	Content string                      `json:"content"`
}

type RecommendationsAnimesData struct {
	Title string                        `json:"title"`
	MalId int                           `json:"mal_id"`
	Url   string                        `json:"url"`
	Image RecommendationsImagesJikanApi `json:"images"`
}

type RecommendationsImagesJikanApi struct {
	JPG struct {
		ImageUrl string `json:"image_url"`
	} `json:"jpg"`
}
