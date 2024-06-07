package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Name    string `json:"Name"`
	Age     int    `json:"Age`
	IsAdult bool   `json:"isAdult`
}

func main() {
	fmt.Println("i m Learing JSON")
	person := Person{Name: "Sankalp", Age: 21, IsAdult: true}
	//fmt.Println("Person Data is :", person)

	// convert person into JSON Encoding (Marshalling)
	jsondata,err:=json.Marshal(person)
	if err!=nil{
		fmt.Println("Got an Error during Marshling",err)
		return
	}
	fmt.Println("person data is :",string(jsondata))

	var persondata Person
	err =json.Unmarshal(jsondata,&persondata)
	if err!=nil{
		fmt.Println("Error Handling ",err)
		return
	}
	fmt.Println("perosn data after Unmarshling is  :",persondata)
}
