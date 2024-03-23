package controllers

import (
	"mygram_finalprojectgo/database"
	"mygram_finalprojectgo/models"
	"net/http"
	"strconv"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

//mebuat data comment
	func CreateComment(c *gin.Context) {
	db := database.GetDB()
	userData := c.MustGet("userData").(jwt.MapClaims)
	userID := uint(userData["id"].(float64))

	photoID, err := strconv.Atoi(c.Param("photoId"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": "Invalid photo ID",
		})
		return
	}

	photo := models.Photo{}
	if err := db.Where("id = ?", photoID).First(&photo).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error":   "Not Found",
			"message": "Photo not found",
		})
		return
	}

	comment := models.Comment{}
	if err := c.BindJSON(&comment); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": "Invalid comment data",
		})
		return
	}
	comment.UserID = userID
	comment.PhotoID = uint(photoID)

	if err := db.Create(&comment).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "Internal Server Error",
			"message": "Failed to add comment",
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"id": comment.ID, 
		"message": comment.Message, 
		"photo_id": comment.PhotoID, 
		"user_id": comment.UserID, 
		"created_at": comment.CreatedAt, 
	})
}

// Update Comment salah
func PutComment(c *gin.Context) {
	db := database.GetDB()
	userData := c.MustGet("userData").(jwt.MapClaims)
	userID := uint(userData["id"].(float64))

	photoID, err := strconv.Atoi(c.Param("photoId"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": "Invalid photo ID",
		})
		return
	}

	commentID, err := strconv.Atoi(c.Param("commentId"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": "Invalid comment ID",
		})
		return
	}

	var comment models.Comment
	if err := db.Where("id = ? AND photo_id = ?", commentID, photoID).First(&comment).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error":   "Not Found",
			"message": "Comment not found for the given photo",
		})
		return
	}

	if comment.UserID != userID { // Ensure comment belongs to the user
		c.JSON(http.StatusUnauthorized, gin.H{
			"error":   "Unauthorized",
			"message": "You are not authorized to update this comment",
		})
		return
	}

	var updatedComment models.Comment
	if err := c.BindJSON(&updatedComment); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": "Invalid comment data",
		})
		return
	}

	// Update comment fields
	comment.Message = updatedComment.Message

	if err := db.Save(&comment).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "Internal Server Error",
			"message": "Failed to update comment",
		})
		return
	}

	c.JSON(http.StatusOK, comment)
}
