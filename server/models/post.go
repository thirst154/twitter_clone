package models

type Post struct {
	ID            uint   `json:"id" gorm:"primary_key"`
	Text          string `json:"text"`
	DatePublished uint   `json:"date_published"`
	AuthorID      uint   `json:"author_id"`
}
