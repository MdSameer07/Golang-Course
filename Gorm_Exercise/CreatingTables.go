package main

import (
	_ "fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
)

type Employee struct{
	gorm.Model
	EmpName string `gorm:"not null"`
	DepartmentId uint  `gorm:"not null"`
	Department Department  `gorm:"not null"`
	ManagerId uint  `gorm:"not null"`
	Manager *Employee
}

type Department struct{
	gorm.Model
	Name string  `gorm:"not null"`
}

func main2(){
	db,err := gorm.Open("postgres","user=sameer password=******** dbname=exercise host=localhost port=5432 sslmode=disable")  
	if err!=nil{
		panic(err.Error())
	}
	defer db.Close()

	db.AutoMigrate(&Employee{},&Department{})
	db.Model(&Employee{}).AddForeignKey("department_id","departments(id)","CASCADE","CASCADE")
	println("done")
}