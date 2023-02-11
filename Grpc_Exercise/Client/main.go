package main

import (
	"context"
	"log"
	_ "github.com/lib/pq"
	"google.golang.org/grpc"
	"news.com/events/Grpc_Exercise/employee"
)

func main() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
	defer conn.Close()

	client := employee.NewEmployeeServiceClient(conn)

	//Creating(C)

	resp1, err := client.CreateEmployee(context.Background(), &employee.CreateEmployeeRequest{EmployeeName: "Sameer",ManagerName: "Teja",Department: "Coding"})
	if err != nil {
		log.Fatalf("Failed to create employee: %v", err)
	}

	log.Printf("Employee created successfully with ID: %s", resp1.GetMessage())

	//Reading(R)

	resp2,err := client.ReadEmployee(context.Background(),&employee.ReadEmployeeRequest{EmployeeId: 2})
	if err!=nil{
		log.Fatalf("Failed to read employee: %v",err)
	}

	log.Println("Employee Details : ",resp2.GetEmployee())

	//Updating(U)

	resp3,err := client.UpdateEmployee(context.Background(),&employee.UpdateEmployeeRequest{EmployeeName: "Sameer",ManagerName: "Sai"})
	if err!=nil{
		log.Fatalf("Failed to Update employee :%v",err)
	}

	log.Printf(resp3.GetMessage())

	//Delete(D)

	resp4,err := client.DeleteEmployee(context.Background(),&employee.DeleteEmployeeRequest{EmployeeName: "Teja"})
	if err!=nil{
		log.Fatalf("Failed to Delete employee :%v",err)
	}

	log.Printf(resp4.GetMessage())

}