package models

type Song struct {
	ID          uint   `json:"id" gorm:"primaryKey"`
	Group       string `json:"group"`
	SongTitle   string `json:"song"`
	ReleaseDate string `json:"releaseDate"`
	Text        string `json:"text"`
	Link        string `json:"link"`
}
