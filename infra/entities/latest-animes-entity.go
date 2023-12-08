package entities

type LatestAnimesEntity struct {
	MalId    int    `json:"mal_id"`
	Url      string `json:"url"`
	Title    string `json:"title"`
	Rank     int    `json:"rank"`
	ImageUrl string `json:"image_url"`
}
