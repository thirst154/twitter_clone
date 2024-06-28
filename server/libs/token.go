package libs

import (
	"example/server/models"

	"github.com/gin-gonic/gin"
)

func GetAuth(c *gin.Context) (string, error) {
	var auth AuthHeader
	// Get Auth
	if err := c.ShouldBindHeader(&auth); err != nil {
		return "", err
	}

	token := auth.Auth

	return token, nil
}

func GetTokenReturnUser(c *gin.Context) (*models.User, error) {
	token, err := GetAuth(c)
	if err != nil {
		return nil, err
	}
	// get user from token
	user, err := models.GetUserFromToken(token)
	if err != nil {
		return nil, err
	}

	return user, nil
}
