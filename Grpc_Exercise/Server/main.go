package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"
	"github.com/jinzhu/gorm"
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
	EmpName string `gorm:"not null;unique"`
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
	db,err := gorm.Open("postgres","user=sameer password=19189149 dbname=exercise host=localhost port=5432 sslmode=disable")  
	if err!=nil{
		return nil,status.Errorf(codes.Aborted,"Connection with Database Failed")
	}
	defer db.Close()

	var dept Department
	db.FirstOrCreate(&dept,Department{Name: req.GetDepartment()})

	var emp Employee
	db.FirstOrCreate(&emp,Employee{EmpName: req.GetManagerName()})

	if emp.DepartmentId != 0 && emp.DepartmentId != dept.ID {
        return nil,status.Errorf(codes.FailedPrecondition,"Employee and Manager Department Should match")
    }
    if emp.DepartmentId == 0 {
        emp.DepartmentId = dept.ID
        db.Save(&emp)
    }
	
	employeeStruct := &Employee{
		EmpName : req.GetEmployeeName(),
		DepartmentId: dept.ID,
		ManagerId: &emp.ID,
	}

	if err := db.Create(employeeStruct).Error; err!=nil{
		return nil,status.Errorf(codes.Internal,"Creation of employee failed")
	}
	resp := &employee.CreateEmployeeResponse{
		Message: fmt.Sprintf("Employee created successfully with Id : %d",employeeStruct.ID),
	}

	return resp,nil
}

func (s *employeeServiceServer) ReadEmployee(ctx context.Context,req *employee.ReadEmployeeRequest) (*employee.ReadEmployeeResponse,error){
	db,err := gorm.Open("postgres","user=sameer password=19189149 dbname=exercise host=localhost port=5432 sslmode=disable")  
	if err!=nil{
		return nil,status.Errorf(codes.Aborted,"Connection with Database Failed")
	}
	defer db.Close()

	var emp Employee
	db.Model(&Employee{}).Where("id=?",req.GetEmployeeId()).Find(&emp)

	if emp==(Employee{}){
		return nil, status.Errorf(codes.NotFound, "Employee with ID %d not found", req.GetEmployeeId())
	}

	resp :=  &employee.ReadEmployeeResponse{
		Employee: &employee.Employee{
			Id: uint32(emp.ID),
			Name: emp.EmpName,
			DepartmentId: uint32(emp.DepartmentId),
		},
	}
	if emp.ManagerId!=nil{
		resp.Employee.ManagerId = uint32(*emp.ManagerId)
	}

	return resp,nil
}

func (s *employeeServiceServer) UpdateEmployee(ctx context.Context,req *employee.UpdateEmployeeRequest) (*employee.UpdateEmployeeResponse,error){
	db,err := gorm.Open("postgres","user=sameer password=19189149 dbname=exercise host=localhost port=5432 sslmode=disable")  
	if err!=nil{
		return nil,status.Errorf(codes.Aborted,"Connection with Database Failed")
	}
	defer db.Close()

	var emp1 Employee
	db.Model(&Employee{}).Where("emp_name=?",req.GetEmployeeName()).Find(&emp1)

	if emp1==(Employee{}){
		return nil,status.Errorf(codes.NotFound,"Employee doesn't exist in the employee table")
	}

	var emp2 Employee

	db.FirstOrCreate(&emp2,Employee{EmpName:req.GetManagerName()})

	if emp2.DepartmentId!=0 && emp1.DepartmentId!=emp2.DepartmentId{
		return nil,status.Errorf(codes.FailedPrecondition,"Employee and Manager Departments should match")
	}

	if emp2.DepartmentId==0{
		emp2.DepartmentId = emp1.DepartmentId
		db.Save(&emp2)
	}

	emp1.ManagerId = &emp2.ID
	db.Save(&emp1)

	resp := &employee.UpdateEmployeeResponse{
		Message: fmt.Sprintf("Employee with Id : %d updated successfully",emp1.ID),
	}

	return resp,nil
}


func (s *employeeServiceServer) DeleteEmployee(ctx context.Context,req *employee.DeleteEmployeeRequest) (*employee.DeleteEmployeeResponse,error){
	db,err := gorm.Open("postgres","user=sameer password=19189149 dbname=exercise host=localhost port=5432 sslmode=disable")  
	if err!=nil{
		return nil,status.Errorf(codes.Aborted,"Connection with Database Failed")
	}
	defer db.Close()

	var emp Employee
	db.Model(&Employee{}).Where("emp_name=?",req.GetEmployeeName()).Find(&emp)

	db.Model(&Employee{}).Where("manager_id = ?", emp.ID).Updates(map[string]interface{}{"manager_id": nil})
	
	db.Delete(&emp)

	resp := &employee.DeleteEmployeeResponse{
		Message: fmt.Sprintf("Employee with name :%s deleted successfully",req.GetEmployeeName()),
	}

	return resp,nil
}

func main(){
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	employee.RegisterEmployeeServiceServer(s,&employeeServiceServer{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}