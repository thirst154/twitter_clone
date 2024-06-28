package libs

import (
	"example/server/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AuthHeader struct {
	Auth string `header:"X-Authorization"`
}

func IsAuthenticated() gin.HandlerFunc {
	return func(c *gin.Context) {

		token, err := GetAuth(c)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Token Required"})
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
