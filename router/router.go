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

		photoRouter.POST("/:photoId/comments", controllers.CreateComment)
		photoRouter.PUT("/:photoId/comments/:commentId", midleware.CommentAuthorization(), controllers.PutComment)
		photoRouter.DELETE("/:photoId/comments/:commentId", midleware.CommentAuthorization(), controllers.DeleteComment)
		photoRouter.GET("/:photoId/comments", midleware.PhotoAuthorization(),controllers.GetComment)

		photoRouter.PUT("/:photoId", midleware.PhotoAuthorization(), controllers.UpdatePhoto)
		photoRouter.DELETE("/:photoId", midleware.PhotoAuthorization(), controllers.DeletePhoto)
		photoRouter.GET("/user/:userID", midleware.PhotoAuthorization(), controllers.GetPhoto) //masih belum muncul
	}

	socialMediaRouter := r.Group("/socialmedias")
	{
		socialMediaRouter.Use(midleware.Authentication())
		socialMediaRouter.POST("/", controllers.PostSocialMedia)
		socialMediaRouter.PUT("/:socialMediaId", controllers.PutSocialMedia)
		socialMediaRouter.DELETE("/:socialMediaId", controllers.DeleteSocialMedia)
		socialMediaRouter.GET("/", midleware.Authentication(),controllers.GetSocialMedia)
	}

	// commentRouter := r.Group("/comments") //belum masih stuck
	// {
	// 	commentRouter.Use(midleware.Authentication())
	// 	commentRouter.POST("/", controllers.CreateComment)
	// }

	return r
}