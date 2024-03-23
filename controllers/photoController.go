package controllers

import (
	"mygram_finalprojectgo/database"
	"mygram_finalprojectgo/helpers"
	"mygram_finalprojectgo/models"
	"net/http"
	"strconv"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

//membuat data Photo
func CreatePhoto(c *gin.Context) {
	db := database.GetDB()
	userData := c.MustGet("userData").(jwt.MapClaims)
	contentType := helpers.GetContentType(c)

	Photo := models.Photo{}
	userID := uint(userData["id"].(float64))

	if contentType == appJSON {
		c.ShouldBindJSON(&Photo)
	} else {
		c.ShouldBind(&Photo)
	}

	Photo.UserID = userID

	err := db.Debug().Create(&Photo).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err": "Bad Request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, Photo)
}

//membuat Update Photo

func UpdatePhoto(c *gin.Context) {
	db := database.GetDB()
	userData := c.MustGet("userData").(jwt.MapClaims)
	contentType := helpers.GetContentType(c)
	Photo := models.Photo{}

	photoId, _ := strconv.Atoi(c.Param("photoId"))
	userId := uint(userData["id"].(float64))

	if contentType == appJSON {
		c.ShouldBindJSON(&Photo)
	} else {
		c.ShouldBind(&Photo)
	}

	Photo.UserID = userId
	Photo.ID = uint(photoId)

	err := db.Model(&Photo).Where("id = ?", photoId).Updates(models.Photo{Title: Photo.Title, Caption: Photo.Caption, PhotoUrl: Photo.PhotoUrl}).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err": "Bad Request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, Photo)
}

//Delete Photo

func DeletePhoto(c *gin.Context) {
	db := database.GetDB()
	userData := c.MustGet("userData").(jwt.MapClaims)
	photoID, err := strconv.Atoi(c.Param("photoId"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid photo ID",
		})
		return
	}
	userID := uint(userData["id"].(float64))

	Photo := models.Photo{}
	if err := db.First(&Photo, photoID).Error; err != nil{
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Photo not Found",
		})
		return
	}

	if Photo.UserID != userID {
		c.JSON(http.StatusForbidden, gin.H{
			"error": "You do not have permission to delete",
		})
		return
	}

	if err := db.Delete(&Photo).Error; err != nil{
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to delete photo",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Photo has been deleted",
	})
}

//Get data photo

func GetPhoto(c *gin.Context) {
    db := database.GetDB()
    userID := c.Param("userID")

    var photos []models.Photo
    if err := db.Where("user_id = ?", userID).Find(&photos).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{
            "error":   "Internal Server Error",
            "message": "Failed to retrieve photos",
        })
        return
    }

    c.JSON(http.StatusOK, photos)
}