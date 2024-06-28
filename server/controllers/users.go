package controllers

import (
	"example/server/libs"
	"example/server/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

type NewUserInput struct {
	FirstName string `json:"first_name" binding:"required"`
	LastName  string `json:"last_name" binding:"required"`
	Username  string `json:"username" binding:"required"`
	Password  string `json:"password" binding:"required"`
}

type LoginInput struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type LogoutInput struct {
	ID uint `json:"ID" binding:"required"`
}

// POST /users
// add new users
func CreateUser(c *gin.Context) {
	var input NewUserInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if conflict, _ := models.GetUserFromUsername(input.Username); conflict != nil {
		c.JSON(http.StatusConflict, gin.H{"error": "Username already taken"})
		return
	}

	salt := libs.GenerateRandomSalt(libs.SaltSize)

	user := models.User{
		FirstName: input.FirstName,
		LastName:  input.LastName,
		Username:  input.Username,
		Password:  libs.HashPassword(input.Password, salt),
		Salt:      string(salt),
	}

	models.DB.Create(&user)

	c.JSON(http.StatusOK, gin.H{"ID": user.ID})
}

// GET /users/:id
// Gete user details
func GetDetails(c *gin.Context) {
	var user models.User

	if err := models.DB.Where("id = ?", c.Param("id")).First(&user).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record Not Found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"ID": user.ID, "first_name": user.FirstName, "last_name": user.LastName, "username": user.Username})

}

// POST /login
// Log in a user
func Login(c *gin.Context) {
	var input LoginInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := models.GetUserFromUsername(input.Username)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "cannot find User"})
		return
	}

	if !libs.CheckPassword(user.Password, input.Password, []byte(user.Salt)) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Password Incorrect"})
		return
	}

	// Try Get Token
	if user.SessionToken == "0" || user.SessionToken == "" {
		token, err := models.SetToken(user.ID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err})
			return
		}

		c.JSON(http.StatusOK, gin.H{"id": user.ID, "token": token})
		return

	}

	c.JSON(http.StatusOK, gin.H{"id": user.ID, "token": user.SessionToken})
}

// POST /logout
// log out a user
func Logout(c *gin.Context) {

	// Get Auth
	user, err := libs.GetTokenReturnUser(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Cannot Find User"})
		return
	}

	_, err = models.RemoveToken(user.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed To remove token from user"})
		return
	}

	c.JSON(http.StatusAccepted, gin.H{"id": user.ID})
}
