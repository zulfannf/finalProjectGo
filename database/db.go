package database

import (
	"fmt"
	"log"
	"mygram_finalprojectgo/models"
	// "os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	host = ""
	user = ""
	password = ""
	port = ""
	dbname = ""
	db	*gorm.DB
	err	error

	// host = os.Getenv("PGHOST")
	// user = os.Getenv("PGUSER")
	// password = os.Getenv("PGPASSWORD")
	// port = os.Getenv("PGPORT")
	// dbname = os.Getenv("PGDATABASE")
	// db	*gorm.DB
	// err	error
)

func StartDB(){
	config := fmt.Sprintf("host=%s user=%s password=%s port=%s dbname=%s sslmode=disable", host, user, password, port, dbname)
	dsn := config
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("error connecting to database :", err)
	}

	fmt.Println("Sukses koneksi ke database :", err)
	db.Debug().AutoMigrate(models.User{}, models.Photo{}, models.Comment{}, models.SocialMedia{})
}	
	
func GetDB() *gorm.DB{
	return db
}
