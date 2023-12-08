package entities

type RecommendationsAnimesEntity struct {
	MalId    int    `json:"mal_id"`
	Title    string `json:"Title"`
	Content  string `json:"content"`
	Url      string `json:"url"`
	ImageUrl string `json:"image_url"`
}
