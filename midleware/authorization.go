package midleware

import (
	"mygram_finalprojectgo/database"
	"mygram_finalprojectgo/models"
	"net/http"
	"strconv"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func PhotoAuthorization() gin.HandlerFunc {
	return func(c *gin.Context) {
		db := database.GetDB()
		photoId, err := strconv.Atoi(c.Param("photoId"))
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"error": "Bad Request",
				"message": "invalid parameter",
			})
			return
		}
		userData := c.MustGet("userData").(jwt.MapClaims)
		userId := uint(userData["id"].(float64))
		Photo := models.Photo{}

		err = db.Select("user_id").First(&Photo, uint(photoId)).Error

		if err != nil {
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
				"error": "Data Not Found",
				"message": "data doesn't exist",
			})
			return
		}

		if Photo.UserID != userId {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "Unauthorized",
				"message": "you are not allowed to access this data",
			})
			return
		}

		c.Next()
	}
}

func CommentAuthorization() gin.HandlerFunc {
	return func (c *gin.Context)  {
		db := database.GetDB()
		userData := c.MustGet("userData").(jwt.MapClaims)
		userID := uint(userData["id"].(float64))

		commentID, err := strconv.Atoi(c.Param("commentId"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "Bad Request",
				"message": "Invalid comment ID",
			})
			c.Abort()
			return
		}

		comment := models.Comment{}
		if err := db.First(&comment, commentID).Error; err != nil {
			c.JSON(http.StatusNotFound, gin.H{
				"error": "Not Found",
				"message": "Comment not Found",
			})
			c.Abort()
			return
		}

		if comment.UserID != userID {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "Unauthorized",
				"message": "You are not authorized to access this comment",
			})
			c.Abort()
			return
		}

		c.Next()
	}
}