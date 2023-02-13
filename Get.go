package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type User2 struct{
	Page int
	Total int
}

func main2(){
	m := User2{}
	resp,err := http.Get("https://reqres.in/api/users?page=2")
	if err!=nil{
		fmt.Println(err)
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
	fmt.Println(m.Page)
	fmt.Println(m.Total)
}