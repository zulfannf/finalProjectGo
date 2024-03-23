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
		userRouter.PUT("/:userID", controllers.UpdateUser) //masih eror data belum ke rubah
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

	return r
}