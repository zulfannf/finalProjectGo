package main

import (
	"mygram_finalprojectgo/database"
	"mygram_finalprojectgo/router"
	"os"
)

// func main(){
// 	database.StartDB()
// 	r := router.StartApp()
// 	r.Run(":8080")
// }

func main() {
	database.StartDB()

	var PORT = os.Getenv("PORT")

	router.StarServer().Run(":" + PORT)
}
