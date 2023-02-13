package main

import (
	"context"
	"log"
	_ "github.com/lib/pq"
	"google.golang.org/grpc"
	"news.com/events/Grpc_Exercise/employee"
)

func main() {
	Conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
		return 
	}
	defer Conn.Close()

	client := employee.NewEmployeeServiceClient(Conn)

	//Creating(C)

	resp1, err := client.CreateEmployee(context.Background(), &employee.CreateEmployeeRequest{EmployeeName: "Sameer",EmployeeEmail: "sameer1@gmail.com",ManagerName: "Sai",ManagerEmail: "sai1@gmail.com",Department: "Coding"})
	if err != nil {
		log.Fatalf("Failed to create employee: %v", err)
		return 
	}

	log.Println("Employee Created Successfully : ", resp1.GetEmployee())

	//Reading(R)

	resp2,err := client.ReadEmployee(context.Background(),&employee.ReadEmployeeRequest{EmployeeId: 2})
	if err!=nil{
		log.Fatalf("Failed to read employee: %v",err)
		return 
	}

	log.Println("Employee Details : ",resp2.GetEmployee())

	//Updating(U)

	resp3,err := client.UpdateEmployee(context.Background(),&employee.UpdateEmployeeRequest{EmployeeEmail: "sameer1@gmail.com",ManagerName: "Teja",ManagerEmail: "teja1@gmail.com"})
	if err!=nil{
		log.Fatalf("Failed to Update employee :%v",err)
		return 
	}

	log.Println("Updated Employee : ",resp3.GetEmployee())

	//Delete(D)

	resp4,err := client.DeleteEmployee(context.Background(),&employee.DeleteEmployeeRequest{EmployeeEmail: "teja1@gmail.com"})
	if err!=nil{
		log.Fatalf("Failed to Delete employee :%v",err)
		return 
	}

	log.Printf(resp4.GetMessage())

}