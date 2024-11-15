package main

import (
	_ "fmt"
	"log"
	"os"
	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

// haha this functions doesn't look good
func main1(){
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
		return
	}

	Db_details := os.Getenv("DB_Details")
	db,err := gorm.Open("postgres",Db_details)  
	if err!=nil{
		panic(err.Error())
	}
	defer db.Close()
	dbase := db.DB()  
	defer dbase.Close()
	err = dbase.Ping()
	if err!=nil{
		panic(err.Error())
	}
	println("Connection to database was successfully established")
}