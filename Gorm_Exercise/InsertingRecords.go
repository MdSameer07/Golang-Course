package main

import (
	"fmt"
	"log"
	"os"
	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

//Marketing,Finance,Sales,Purchasing,Operations

func main3(){
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

	if err = db.AutoMigrate(&Employee{},&Department{}).Error; err!=nil{
		fmt.Println(err)
		return 
	}

	departments := []Department{
		{Name:"Marketing"},
		{Name:"Finance"},
		{Name:"Sales"},
		{Name:"Purchasing"},
		{Name:"Operations"},
	}

	for _,department := range departments{
		if err = db.Create(&department).Error; err!=nil{
			fmt.Println(err)
			return
		}
	}

	var dept Department
	if err = db.FirstOrCreate(&dept,Department{Name: "Coding"}).Error; err!=nil{
		fmt.Println(err)
		return 
	}

	var emp Employee
	if err = db.FirstOrCreate(&emp,Employee{EmpName: "Sai",EmpEmail: "sai1@gmail.com"}).Error; err!=nil{
		fmt.Println(err)
		return 
	}

	if emp.DepartmentId != 0 && emp.DepartmentId != dept.ID {
        fmt.Println("Manager's department is different from employee's department. Cannot create new employee.")
        return
    }
    if emp.DepartmentId == 0 {
        emp.DepartmentId = dept.ID
        if err = db.Save(&emp).Error;err!=nil{
			fmt.Println(err)
			return 
		}
    }

	empl := &Employee{
		EmpName : "Teja",
		EmpEmail: "teja1@gmail.com",
		DepartmentId: dept.ID,
		ManagerId: &emp.ID,
	}

	if err = db.Create(&empl).Error; err!=nil{
		fmt.Println(err)
		return 
	}

	var dept2 Department
	if err = db.FirstOrCreate(&dept2,Department{Name: "Testing"}).Error; err!=nil{
		fmt.Println(err)
		return 
	}

	var emp2 Employee
	if err = db.FirstOrCreate(&emp2,Employee{EmpName: "Sai",EmpEmail: "sai1@gmail.com"}).Error; err!=nil{
		fmt.Println(err)
		return
	}

	if emp2.DepartmentId != 0 && emp2.DepartmentId != dept2.ID {
        fmt.Println("Manager's department is different from employee's department. Cannot create new employee.")
        return
    }
    if emp2.DepartmentId == 0 {
        emp2.DepartmentId = dept2.ID
        if err = db.Save(&emp2).Error; err!=nil{
			fmt.Println(err)
			return 
		}
    }

	employ := &Employee{
		EmpName: "Naga",
		EmpEmail: "naga1@gmail.com",
		DepartmentId: dept2.ID,
		ManagerId: &emp2.ID,
	}

	if err = db.Create(&employ).Error; err!=nil{
		fmt.Println(err)
		return
	}

	// employees := []Employee{
	// 	{EmpName: "Sameer"},
	// 	{EmpName: "Tarun"},
	// 	{EmpName: "Adithya"},
	// 	{EmpName: "Shreekar"},
	// 	{EmpName: "Madhu"},
	// 	{EmpName: "Nishanth"},
	// 	{EmpName: "Thrilok"},
	// 	{EmpName: "Revi"},
	// 	{EmpName: "Random"},
	// }

	// for _,employee := range employees{
	// 	db.Create(&employee)
	// }

	// Update(db,"Sameer","Shreekar","Operations")
	// Update(db,"Tarun","Sameer","Marketing")
	// Update(db,"Adithya","Revi","Finance")
	// Update(db,"Shreekar","Random","Purchasing")
	// Update(db,"Madhu","Sameer","Sales")
	// Update(db,"Nishanth","Madhu","Finance")
	// Update(db,"Thrilok","Adithya","Marketing")
	// Update(db,"Revi","Shreekar","Operations")
	// Update(db,"Random","Shreekar","Purchasing")
	
	// println("done")
}

func Update(db *gorm.DB,Name,M_Name,Dept string){
	if err := db.Model(Employee{}).Where("emp_name = ?",Name).Updates(map[string]interface{}{
		"ManagerId": db.Model(Employee{}).Where("emp_name = ?",M_Name).Select("id").SubQuery(),
		"department_id": db.Model(Department{}).Where("name=?",Dept).Select("id").SubQuery(),
	}).Error;err!=nil{
		fmt.Println(err)
		return 
	}
}

func (e *Employee) AfterCreate() error{
	fmt.Println("User created successfully")
	return nil
}