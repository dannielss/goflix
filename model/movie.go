package model

type Movie struct {
	Id          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
}
type PayloadMovie struct {
	Movie      Movie `json:"movie"`
	CategoryId int   `json:"category_id"`
}

type MovieWithCategory struct {
	Id          int64  `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Category    string `json:"category"`
}
