package libs

import (
	"example/server/models"
	"time"
)

type ReturnUser struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Username  string `json:"username"`
}

type ReturnPost struct {
	ID            uint       `json:"id"`
	Text          string     `json:"text"`
	DatePublished time.Time  `json:"date_published"`
	User          ReturnUser `json:"user"`
}

func FormatPost(post models.Post) (ReturnPost, error) {
	user, err := models.GetUserFromID(post.AuthorID)
	if err != nil {
		return ReturnPost{}, err
	}

	returnUser := ReturnUser{
		FirstName: user.FirstName,
		LastName:  user.FirstName,
		Username:  user.Username,
	}

	returnPost := ReturnPost{
		ID:            post.ID,
		Text:          post.Text,
		DatePublished: post.DatePublished,
		User:          returnUser,
	}

	return returnPost, err

}
