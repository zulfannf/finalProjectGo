package controllers

import (
	"mygram_finalprojectgo/database"
	"mygram_finalprojectgo/models"
	"net/http"
	"strconv"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func PostSocialMedia(c *gin.Context) {
	db := database.GetDB()
	userData := c.MustGet("userData").(jwt.MapClaims)
	userID := uint(userData["id"].(float64))

	var socialMedia models.SocialMedia
	if err := c.ShouldBindJSON(&socialMedia); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": "Invalid request body",
		})
		return
	}

	socialMedia.UserID = userID

	if err := db.Create(&socialMedia).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "Internal Server Error",
			"message": "Failed to create social media",
		})
		return
	}

	c.JSON(http.StatusCreated, socialMedia)
}

//Update
func PutSocialMedia(c *gin.Context) {
	db := database.GetDB()
	userData := c.MustGet("userData").(jwt.MapClaims)
	userID := uint(userData["id"].(float64))

	socialMediaID, err := strconv.Atoi(c.Param("socialMediaId"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": "Invalid social media ID",
		})
		return
	}

	var socialMedia models.SocialMedia
	if err := db.First(&socialMedia, socialMediaID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error":   "Not Found",
			"message": "Social media not found",
		})
		return
	}

	if socialMedia.UserID != userID {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error":   "Unauthorized",
			"message": "You are not authorized to update this social media",
		})
		return
	}

	var updatedSocialMedia models.SocialMedia
	if err := c.BindJSON(&updatedSocialMedia); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": "Invalid social media data",
		})
		return
	}

	// Update social media fields
	socialMedia.Name = updatedSocialMedia.Name
	socialMedia.SocialMediaUrl = updatedSocialMedia.SocialMediaUrl

	if err := db.Save(&socialMedia).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "Internal Server Error",
			"message": "Failed to update social media",
		})
		return
	}

	c.JSON(http.StatusOK, socialMedia)
}

//Delete Social Media
func DeleteSocialMedia(c *gin.Context) {
	db := database.GetDB()

	// Mendapatkan data pengguna dari middleware
	userData := c.MustGet("userData").(jwt.MapClaims)
	userID := uint(userData["id"].(float64))

	// Mendapatkan ID media sosial dari parameter URL
	socialMediaID, err := strconv.Atoi(c.Param("socialMediaId"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": "Invalid social media ID",
		})
		return
	}

	// Mencari media sosial dengan ID yang diberikan dari database
	var socialMedia models.SocialMedia
	if err := db.First(&socialMedia, socialMediaID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error":   "Not Found",
			"message": "Social media not found",
		})
		return
	}

	// Memeriksa apakah pengguna memiliki izin untuk menghapus media sosial
	if socialMedia.UserID != userID {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error":   "Unauthorized",
			"message": "You are not authorized to delete this social media",
		})
		return
	}

	// Menghapus media sosial dari database
	if err := db.Delete(&socialMedia).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "Internal Server Error",
			"message": "Failed to delete social media",
		})
		return
	}

	// Mengirim respons bahwa media sosial berhasil dihapus
	c.JSON(http.StatusOK, gin.H{
		"message": "Social media deleted successfully",
	})
}

//Get Social Media

func GetSocialMedia(c *gin.Context) {
	db := database.GetDB()

	// Mendapatkan data pengguna dari middleware
	userData := c.MustGet("userData").(jwt.MapClaims)
	userID := uint(userData["id"].(float64))

	// Mencari semua media sosial milik pengguna dari database
	var socialMedia []models.SocialMedia
	if err := db.Where("user_id = ?", userID).Find(&socialMedia).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "Internal Server Error",
			"message": "Failed to fetch social media",
		})
		return
	}

	// Mengirim respons dengan data media sosial
	c.JSON(http.StatusOK, socialMedia)
}