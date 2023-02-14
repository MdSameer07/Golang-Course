package main

import (
	"fmt"
	_ "fmt"
	"log"
	"os"
	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

type Employee struct{
	gorm.Model
	EmpName string `gorm:"not null"`
	EmpEmail string `gorm:"not null;unique"`
	DepartmentId uint  `gorm:"not null"`
	Department Department  `gorm:"not null"`
	ManagerId *uint  
	Manager *Employee
}

type Department struct{
	gorm.Model
	Name string  `gorm:"not null;unique"`
}

func main2(){
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

	if err = db.DropTable(&Employee{}).Error; err!=nil{
		fmt.Println(err)
		return 
	}
	if err = db.DropTable(&Department{}).Error; err!=nil{
		fmt.Println(err)
		return 
	}
	db.AutoMigrate(&Employee{},&Department{})
	if err = db.Model(&Employee{}).AddForeignKey("department_id","departments(id)","CASCADE","CASCADE").Error;err!=nil{
		fmt.Println(err)
		return
	}
	println("done")
}