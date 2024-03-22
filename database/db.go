package main

import (
	"fmt"
	"log"
	"mygram_finalprojectgo/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	host = "localhost"
	user = "postgres"
	password = "akiyama23"
	port = "5432"
	dbname = "postgres"
	db	*gorm.DB
	err	error
)

func main(){
	config := fmt.Sprintf("host=%s user=%s password=%s port=%s dbname=%s sslmode=disable", host, user, password, port, dbname)
	dsn := config
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("error connecting to database :", err)
	}

	fmt.Println("Sukses koneksi ke database :", err)
	db.Debug().AutoMigrate(models.User{}, models.Photo{})
}	
	
func GetDB() *gorm.DB{
	return db
}
