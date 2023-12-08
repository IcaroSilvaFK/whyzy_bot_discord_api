package entities_external

type LatestAnimesEntityJikanApi struct {
	Data []LatestAnimesDataJikanApi `json:"data"`
}

type LatestAnimesDataJikanApi struct {
	MalId  int                        `json:"mal_id"`
	Url    string                     `json:"url"`
	Images LatestAnimesImagesJikanApi `json:"images"`
	Title  string                     `json:"title"`
	Rank   int                        `json:"rank"`
}

type LatestAnimesImagesJikanApi struct {
	JPG  ImagesJikanApi `json:"jpg"`
	WEBP ImagesJikanApi `json:"webp"`
}

type ImagesJikanApi struct {
	ImageUrl      string `json:"image_url"`
	SmallImageUrl string `json:"small_image_url"`
	LargeImageUrl string `json:"large_image_url"`
}
