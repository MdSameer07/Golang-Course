package main

import (
	_ "fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
)

func main1(){
	db,err := gorm.Open("postgres","user=sameer password=19189149 dbname=exercise host=localhost port=5432 sslmode=disable")  
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