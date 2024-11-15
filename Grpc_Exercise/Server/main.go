package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"
	"os"
	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"news.com/events/Grpc_Exercise/employee"
)

var (
	port = flag.Int("port", 50051, "The server port")
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

type employeeServiceServer struct{
	employee.UnimplementedEmployeeServiceServer
	db *gorm.DB
}

func (s *employeeServiceServer) CreateEmployee(ctx context.Context,req *employee.CreateEmployeeRequest) (*employee.CreateEmployeeResponse,error){

	var dept Department
	if err1 := s.db.FirstOrCreate(&dept,Department{Name: req.GetDepartment()}).Error; err1!=nil{
		return nil,fmt.Errorf("Error occured during reading or creating a department")
	}

	var man Employee
	if err2 := s.db.FirstOrCreate(&man,Employee{EmpName: req.GetManagerName(),EmpEmail: req.GetManagerEmail()}).Error; err2!=nil{
		return nil,fmt.Errorf("Error occured during reading or creating manager")
	}

	if man.DepartmentId != 0 && man.DepartmentId != dept.ID {
        return nil,status.Errorf(codes.FailedPrecondition,"Employee and Manager Department Should match")
    }
    if man.DepartmentId == 0 {
        man.DepartmentId = dept.ID
        s.db.Save(&man)
    }
	
	employeeStruct := &Employee{
		EmpName : req.GetEmployeeName(),
		EmpEmail: req.GetEmployeeEmail(),
		DepartmentId: dept.ID,
		ManagerId: &man.ID,
	}

	if err := s.db.Create(employeeStruct).Error; err!=nil{
		return nil,status.Errorf(codes.Internal,"Creation of employee failed")
	}

	var emp Employee
	s.db.Model(&Employee{}).Where("emp_email=?",req.EmployeeEmail).Find(&emp)


	resp := &employee.CreateEmployeeResponse{
		Employee: &employee.Employee{
			Id: uint32(emp.ID),
			Name: emp.EmpName,
			Email: emp.EmpEmail,
			DepartmentId: uint32(emp.DepartmentId),
			Department: &employee.Department{
				Name: emp.Department.Name,
			},
			ManagerId: uint32(*emp.ManagerId),
			Manager: &employee.Employee{
				Name: emp.Manager.EmpName,
			},
		},
	}

	return resp,nil
}

func (s *employeeServiceServer) ReadEmployee(ctx context.Context,req *employee.ReadEmployeeRequest) (*employee.ReadEmployeeResponse,error){

	var emp Employee
	result := s.db.Model(&Employee{}).Where("id=?",req.GetEmployeeId()).Find(&emp)

	if result.Error != nil{
		if result.RecordNotFound()==true{
			return nil,status.Errorf(codes.NotFound,"Employee doesn't exist in the employee table")
		}
	}

	resp :=  &employee.ReadEmployeeResponse{
		Employee: &employee.Employee{
			Id: uint32(emp.ID),
			Name: emp.EmpName,
			Email: emp.EmpEmail,
			DepartmentId: uint32(emp.DepartmentId),
		},
	}
	if emp.ManagerId!=nil{
		resp.Employee.ManagerId = uint32(*emp.ManagerId)
	}

	return resp,nil
}

func (s *employeeServiceServer) UpdateEmployee(ctx context.Context,req *employee.UpdateEmployeeRequest) (*employee.UpdateEmployeeResponse,error){

	var emp1 Employee
	result := s.db.Model(&Employee{}).Where("emp_email=?",req.GetEmployeeEmail()).Find(&emp1)

	if result.Error != nil{
		if result.RecordNotFound()==true{
			return nil,status.Errorf(codes.NotFound,"Employee doesn't exist in the employee table")
		}
	}

	var emp2 Employee

	if err1 := s.db.FirstOrCreate(&emp2,Employee{EmpName:req.GetManagerName(),EmpEmail: req.GetManagerEmail()}); err1!=nil{
		return nil,fmt.Errorf("Error while reading or creating manager")
	}

	if emp2.DepartmentId!=0 && emp1.DepartmentId!=emp2.DepartmentId{
		return nil,status.Errorf(codes.FailedPrecondition,"Employee and Manager Departments should match")
	}

	if emp2.DepartmentId==0{
		emp2.DepartmentId = emp1.DepartmentId
		s.db.Save(&emp2)
	}

	emp1.ManagerId = &emp2.ID
	s.db.Save(&emp1)

	resp := &employee.UpdateEmployeeResponse{
		Employee: &employee.Employee{
			Id: uint32(emp1.ID),
			Name: emp1.EmpName,
			Email: emp1.EmpEmail,
			ManagerId: uint32(*emp1.ManagerId),
			DepartmentId: uint32(emp1.DepartmentId),
		},
	}

	return resp,nil
}


func (s *employeeServiceServer) DeleteEmployee(ctx context.Context,req *employee.DeleteEmployeeRequest) (*employee.DeleteEmployeeResponse,error){

	var emp Employee
	result := s.db.Model(&Employee{}).Where("emp_email=?",req.GetEmployeeEmail()).Find(&emp)

	if result.Error != nil{
		if result.RecordNotFound()==true{
			return nil,status.Errorf(codes.NotFound,"Employee doesn't exist in the employee table")
		}
	}

	if err := s.db.Model(&Employee{}).Where("manager_id = ?", emp.ID).Updates(map[string]interface{}{"manager_id": nil}).Error; err!=nil{
		return nil,fmt.Errorf("Error while making manager_ids nil")
	}
	
	if err2 := s.db.Delete(&emp).Error; err2!=nil{
		return nil,fmt.Errorf("Error while deleting an employee")
	}

	resp := &employee.DeleteEmployeeResponse{
		Message: fmt.Sprintf("Employee with email :%s deleted successfully",emp.EmpEmail),
	}

	return resp,nil
}

func main(){
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
		return 
	}

	err = godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	Db_details := os.Getenv("DB_Details")

	s := grpc.NewServer()
	db,err := gorm.Open("postgres",Db_details)  
	if err!=nil{
		fmt.Printf("Connection with database failed")
		return
	}
	defer db.Close()
	employee.RegisterEmployeeServiceServer(s,&employeeServiceServer{db:db})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
		return 
	}
}