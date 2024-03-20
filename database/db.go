package main

import (
	"fmt"
	"database/sql"

	_"github.com/lib/pq"
)

const (
	host = "localhost"
	port = 5432
	user = "postgres"
	password = "akiyama23"
	dbname = "postgres"
)

var (
	db *sql.DB
	err error
)

func main(){
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	db, err = sql.Open("postgres",psqlInfo)
	if err != nil {
		panic(err)
	}
	
//Cek koneksi ke data base
	err = db.Ping()
	if err != nil{
		panic(err)
	}
	fmt.Println("Success, connected to database!")
}

func Get()*sql.DB{
	return db
}