package models

type User struct {
	ID           uint   `json:"id" gorm:"primary_key"`
	FirstName    string `json:"first_name"`
	LastName     string `json:"last_name"`
	Username     string `json:"user_name"`
	Password     string `json:"password"`
	Salt         string `json:"salt"`
	SessionToken string `json:"session_token"`
}

func GetUserFromUsername(username string) (*User, error) {
	var user User

	if err := DB.Where("username = ?", username).First(&user).Error; err != nil {
		return nil, err
	}

	return &user, nil

}
