package models

type User struct {
	ID           uint   `json:"id" gorm:"primary_key"`
	FirstName    string `json:"first_name"`
	LastName     string `json:"last_name"`
	UserName     string `json:"user_name"`
	Password     string `json:"password"`
	Salt         string `json:"salt"`
	SessionToken string `json:"session_token"`
}
