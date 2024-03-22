package router

import (
	"mygram_finalprojectgo/controllers"
	"mygram_finalprojectgo/midleware"

	"github.com/gin-gonic/gin"
)

func StartApp() *gin.Engine {
	r := gin.Default()

	userRouter := r.Group("/users")
	{
		userRouter.POST("/register", controllers.UserRegister)

		userRouter.POST("/login", controllers.UserLogin)
		userRouter.PUT("/", controllers.UpdateUser) //masih eror
	}

	photoRouter := r.Group("/photos")
	{
		photoRouter.Use(midleware.Authentication())
		photoRouter.POST("/", controllers.CreatePhoto)

		photoRouter.PUT("/:photoId", midleware.PhotoAuthorization(), controllers.UpdatePhoto)
		photoRouter.DELETE("/:photoId", midleware.PhotoAuthorization(), controllers.DeletePhoto)
		photoRouter.GET("/", midleware.PhotoAuthorization(), controllers.GetPhoto) //masih belum muncul
	}

	return r
}