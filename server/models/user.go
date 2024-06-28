package models

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
)

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

func GetUserFromID(ID uint) (*User, error) {
	var user User
	if err := DB.Where("id = ?", ID).First(&user).Error; err != nil {
		return nil, err
	}

	return &user, nil
}

func GetUserFromToken(token string) (*User, error) {
	var user User
	if err := DB.Where("session_token = ?", token).First(&user).Error; err != nil {
		return nil, err
	}

	fmt.Println(user)

	return &user, nil
}

func updateToken(ID uint, token string) (string, error) {
	user, err := GetUserFromID(ID)
	if err != nil {
		return "", err
	}

	DB.Model(&user).Update("session_token", token)

	return token, nil
}

func RemoveToken(ID uint) (string, error) {
	return updateToken(ID, "")
}

func SetToken(ID uint) (string, error) {
	token, err := randomHex(20)
	if err != nil {
		return "", err
	}
	return updateToken(ID, string(token))
}

func randomHex(n int) (string, error) {
	bytes := make([]byte, n)
	if _, err := rand.Read(bytes); err != nil {
		return "", err
	}
	return hex.EncodeToString(bytes), nil
}
