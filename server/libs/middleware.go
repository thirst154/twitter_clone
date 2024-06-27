package libs

import (
	"example/server/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// func IsAuthenticated(token string) (*models.User, error) {
// 	if token == "0" || token == "" {
// 		return nil, nil
// 	}
// 	user, err := models.GetUserFromToken(token)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return user, nil
// }

type AuthHeader struct {
	Auth string `header:"X-Authorization" binding:"required"`
}

func IsAuthenticated() gin.HandlerFunc {
	return func(c *gin.Context) {

		token, err := GetTokenFromHeader(c)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unauthoerized"})
			return
		}

		if token == "0" || token == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unauthoerized"})
			return
		}

		_, err = models.GetUserFromToken(token)

		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Not a user Token"})
			return
		}

		c.Next()

	}
}

func GetTokenFromHeader(c *gin.Context) (string, error) {
	var auth AuthHeader
	// Get Auth
	if err := c.ShouldBindHeader(auth); err != nil {
		return "", err
	}

	return auth.Auth, nil
}
