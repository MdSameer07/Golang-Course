package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
)

func main4(){
	db,err := gorm.Open("postgres","user=sameer password=19189149 dbname=exercise host=localhost port=5432 sslmode=disable")
	if err!=nil{
		panic(err.Error())
	}
	defer db.Close()
	
	var employees1 []Employee

	db.Preload("Department").Preload("Manager").Find(&employees1)

	for _, employee := range employees1 {
		fmt.Println(employee.EmpName, "belongs to", employee.Department.Name, "Team")
	}

	var employees2 []Employee
	db.Preload("Department").Preload("Manager").Where("manager_id IS NOT NULL").Find(&employees2)

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