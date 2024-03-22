package main

import (
	"mygram_finalprojectgo/database"
	"mygram_finalprojectgo/router"
)

func main(){
	database.StartDB()
	r := router.StartApp()
	r.Run(":8080")
}