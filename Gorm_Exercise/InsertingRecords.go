package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
)

//Marketing,Finance,Sales,Purchasing,Operations

func main3(){
	db,err := gorm.Open("postgres","user=sameer password=******** dbname=exercise host=localhost port=5432 sslmode=disable")
	if err!=nil{
		panic(err.Error())
	}
	defer db.Close()

	db.AutoMigrate(&Employee{},&Department{})

	departments := []Department{
		{Name:"Marketing"},
		{Name:"Finance"},
		{Name:"Sales"},
		{Name:"Purchasing"},
		{Name:"Operations"},
	}

	for _,department := range departments{
		db.Create(&department)
	}

	employees := []Employee{
		{EmpName: "Sameer"},
		{EmpName: "Tarun"},
		{EmpName: "Adithya"},
		{EmpName: "Shreekar"},
		{EmpName: "Madhu"},
		{EmpName: "Nishanth"},
		{EmpName: "Thrilok"},
		{EmpName: "Revi"},
		{EmpName: "Random"},
	}

	for _,employee := range employees{
		db.Create(&employee)
	}

	Update(db,"Sameer","Shreekar","Operations")
	Update(db,"Tarun","Sameer","Marketing")
	Update(db,"Adithya","Revi","Finance")
	Update(db,"Shreekar","Random","Purchasing")
	Update(db,"Madhu","Sameer","Sales")
	Update(db,"Nishanth","Madhu","Finance")
	Update(db,"Thrilok","Adithya","Marketing")
	Update(db,"Revi","Shreekar","Operations")
	Update(db,"Random","Shreekar","Purchasing")
	
	println("done")
}

func Update(db *gorm.DB,Name,M_Name,Dept string){
	db.Model(Employee{}).Where("emp_name = ?",Name).Updates(map[string]interface{}{
		"ManagerId": db.Model(Employee{}).Where("emp_name = ?",M_Name).Select("id").SubQuery(),
		"department_id": db.Model(Department{}).Where("name=?",Dept).Select("id").SubQuery(),
	})
}

func (e *Employee) AfterCreate() error{
	fmt.Println("User created successfully")
	return nil
}