package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type User1 struct {
	Name      string
	Job       string
	Id        string
	CreatedAt string
}

func main3() {
	user := User1{Name: "Sameer", Job: "Developer"}
	m := User1{}
	postBody, _ := json.Marshal(user)
	responseBody := bytes.NewBuffer(postBody)
	resp, err := http.Post("https://reqres.in/api/users", "json", responseBody)
	if err != nil {
		log.Fatalf("An Error Occured %v", err)
		return
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
		return
	}
	err1 := json.Unmarshal(body, &m)
	if err1 != nil {
		log.Fatalln(err1)
		return 
	}
	fmt.Println(m.Id)
	fmt.Println(m.CreatedAt)


	user1 := User1{Name: "Sameer", Job: "Developer"}
	m1 := User1{}
	postBody1, _ := json.Marshal(user1)
	responseBody1 := bytes.NewBuffer(postBody1)
	resp1, err2 := http.Post("https://reqres.in/api/users", "application/json", responseBody1)
	if err2 != nil {
		log.Fatalf("An Error Occured %v", err2)
		return 
	}
	defer resp1.Body.Close()
	body1, err2 := ioutil.ReadAll(resp1.Body)
	if err2 != nil {
		log.Fatalln(err2)
		return
	}
	err3 := json.Unmarshal(body1, &m1)
	if err3 != nil {
		log.Fatalln(err3)
		return 
	}
	fmt.Println(m1.Id)
	fmt.Println(m1.CreatedAt)
}