package main

import (
	"fmt"
	"log"
	"os"
	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func main4(){
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
	
	var employees1 []Employee

	if err = db.Preload("Department").Preload("Manager").Find(&employees1).Error; err!=nil{
		fmt.Println(err)
		return 
	}

	for _, employee := range employees1 {
		fmt.Println(employee.EmpName, "belongs to", employee.Department.Name, "Team")
	}

	var employees2 []Employee
	if err = db.Preload("Department").Preload("Manager").Where("manager_id IS NOT NULL").Find(&employees2).Error;err!=nil{
		fmt.Println(err)
		return 
	}

	for _,employee := range employees2{
		fmt.Println(*(&employee.Manager.EmpName), "is manager of ",employee.EmpName)
	}

	rows,err := db.Table("employees e1").Select("e1.emp_name,e2.emp_name").Joins("join employees e2 on e1.manager_id=e2.id").Rows()

	if err!=nil{
		fmt.Println(err.Error())
	}

	for rows.Next(){
		var name,manager_name string
		rows.Scan(&name,&manager_name)
		fmt.Println(manager_name,"is manager of",name)
	}
}