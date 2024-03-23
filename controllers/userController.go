package controllers

import (
	// "go/token"
	"mygram_finalprojectgo/database"
	"mygram_finalprojectgo/helpers"
	"mygram_finalprojectgo/models"
	"net/http"
	"strconv"

	// "strconv"

	// "github.com/dgrijalva/jwt-go"

	"github.com/gin-gonic/gin"
)

var (
	appJSON = "application/json"
)

func UserRegister(c *gin.Context) {
	db := database.GetDB()
	contentType := helpers.GetContentType(c)
	_, _ = db, contentType
	User := models.User{}

	if contentType == appJSON {
		c.ShouldBindJSON(&User)
	} else {
		c.ShouldBind(&User)
	}

	err := db.Debug().Create(&User).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Bad Request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"id": User.ID,
		"email": User.Email,
		"username": User.Username,
		"age": User.Age,
	})
}

func UserLogin(c *gin.Context) {
	db := database.GetDB()
	contentType := helpers.GetContentType(c)
	_, _ = db, contentType
	User := models.User{}
	password := ""
	

	if contentType == appJSON {
		c.ShouldBindJSON(&User)
	} else {
		c.ShouldBind(&User)
	}

	password = User.Password

	err := db.Debug().Where("email = ?", User.Email).Take(&User).Error

	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Unauthorized",
			"message": "invalid email/password",
		})
		return
	}

	comparePass := helpers.ComparePass([]byte(User.Password), []byte(password))

	if !comparePass {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Unauthorized",
			"message": "invalid email/password",
		})
		return
	}

	token := helpers.GenerateToken(User.ID, User.Email)

	c.JSON(http.StatusOK, gin.H{
		"token": token,
	})
}

//Update User
func UpdateUser(c *gin.Context) {
	db := database.GetDB()
	userID, err := strconv.Atoi(c.Param("userID"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid ID",
		})
		return
	}
	
	User := models.User{}
	if err := db.First(&User, userID).Error; err != nil{
		c.JSON(http.StatusNotFound, gin.H{
			"error": "User not Found",
		})
		return
	}

	if err := db.Model(&User).Updates(&User).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to update",
		})
		return
	}
	
	c.JSON(http.StatusOK, gin.H{
		"id" : User.ID,
		"email" : User.Email,
		"username" : User.Username,
		"age"	: User.Age,
		"update_at" : User.UpdatedAt,
	})
}

//delete User
	func DeleteUser(c *gin.Context) {
		db := database.GetDB()
		userID, err := strconv.Atoi(c.Param("userID"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "Invalid userID",
			})
			return
		}

		user := models.User{}
		if err := db.First(&user, userID).Error; err != nil{
			c.JSON(http.StatusNotFound, gin.H{
				"error": "User Not Found",
			})
			return
		}

		if err := db.Delete(&user).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "failed to delete user",
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"message": "User has been deleted",
		})
	}