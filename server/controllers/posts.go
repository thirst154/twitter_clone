package controllers

import (
	"example/server/libs"
	"example/server/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type NewPostInput struct {
	Text string `json:"text" binding:"required"`
}

// POST /post
// Adds new post from user
func AddPost(c *gin.Context) {
	// Get Auth
	user, err := libs.GetTokenReturnUser(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Cannot Find User"})
		return
	}

	// get input from body req
	var input NewPostInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Add new post via DB

	time := time.Now()

	post := models.Post{
		Text:          input.Text,
		DatePublished: time,
		AuthorID:      user.ID,
	}

	models.DB.Create(&post)

	c.JSON(http.StatusCreated, gin.H{"ID": post.ID})
}

// GET /posts/:id
// get a specific post via ID
func GetPost(c *gin.Context) {
	var post models.Post
	// Get ID from req
	// get post from db
	if err := models.DB.Where("id = ?", c.Param("id")).First(&post).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record Not Found"})
		return
	}

	returnPost, err := libs.FormatPost(post)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Server Failure"})
	}

	// return post
	c.JSON(http.StatusOK, gin.H{"data": returnPost})
}

// Get /posts
// return all posts
func GetAllPosts(c *gin.Context) {
	var posts []models.Post

	models.DB.Find(&posts)

	var returnPosts []libs.ReturnPost
	for _, post := range posts {
		newPost, err := libs.FormatPost(post)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Server Failure"})
		}
		returnPosts = append(
			returnPosts,
			newPost,
		)
	}

	c.JSON(http.StatusOK, gin.H{"data": returnPosts})
}

// PATCH /posts/:id
// Updates a post via ID
func UpdatePost(c *gin.Context) {
	user, err := libs.GetTokenReturnUser(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Cannot Find User"})
		return
	}

	// get post from id
	// get author id from post
	var post models.Post
	if err := models.DB.Where("id = ?", c.Param("id")).First(&post).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Cannot find post"})
		return
	}

	//ensure id's match
	if post.AuthorID != user.ID {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized to edit this post"})
		return
	}
	// get new post info
	var input NewPostInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// update post
	models.DB.Model(&post).Updates(input)

	// return new post
	c.JSON(http.StatusOK, gin.H{"data": post})

}

// DELETE posts/:id
// Delete a post via ID
func DeletePost(c *gin.Context) {
	user, err := libs.GetTokenReturnUser(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Cannot Find User"})
		return
	}

	// get post from id
	// get author id from post
	var post models.Post
	if err := models.DB.Where("id = ?", c.Param("id")).First(&post).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Cannot find post"})
		return
	}

	//ensure id's match
	if post.AuthorID != user.ID {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized to edit this post"})
		return
	}

	// delete post
	models.DB.Delete(&post)

	// return 200
	c.JSON(http.StatusOK, gin.H{"data": true})
}
