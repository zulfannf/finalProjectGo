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
		userRouter.PUT("/:userID", midleware.Authentication(),controllers.UpdateUser) 
		userRouter.DELETE("/:userID", midleware.Authentication(), controllers.DeleteUser)
	}

	photoRouter := r.Group("/photos")
	{
		photoRouter.Use(midleware.Authentication())
		photoRouter.POST("/", controllers.CreatePhoto)

		photoRouter.PUT("/:photoId", midleware.PhotoAuthorization(), controllers.UpdatePhoto)
		photoRouter.DELETE("/:photoId", midleware.PhotoAuthorization(), controllers.DeletePhoto)
		photoRouter.GET("/", midleware.PhotoAuthorization(), controllers.GetPhoto) //masih belum muncul
	}

	commentRouter := r.Group("/comments")
	{
		commentRouter.Use(midleware.Authentication())
		commentRouter.POST("/", controllers.CreateComment)
	}

	return r
}