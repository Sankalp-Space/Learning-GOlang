package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)

type Todo struct {
	UserID    int    `json:"userId"`
	ID        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

func performGetReq() {
	res, err := http.Get("https://jsonplaceholder.typicode.com/todos/1")
	if err != nil {
		fmt.Println("Getting Error :", err)
		return
	}
	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		fmt.Println("Error in getting Response :", res.Status)
	}
	/* data,err:=io.ReadAll(res.Body)
	if err!=nil{
		fmt.Println("Error reading :",err)
		return
	}
	fmt.Println("Data :",string(data)) */

	var todo Todo
	err = json.NewDecoder(res.Body).Decode(&todo)
	if err != nil {
		fmt.Println("Error decoding ", err)
		return
	}
	fmt.Println("Todo :", todo)
	fmt.Println("Title Response :", todo.Title)
	fmt.Println("Completed Response ;", todo.Completed)

}
func performPostReq() {
	todo := Todo{
		UserID:    26,
		Title:     "Sarthak Shrivastav",
		Completed: true,
	}
	// COnvert the Todo struct to JSON
	jsonData, err := json.Marshal(todo)
	if err != nil {
		fmt.Println("Error MArshling :", err)
		return
	}
	//Convert JSON data to string
	jsonString := string(jsonData)

	//convert string data to reader
	jsonReader := strings.NewReader(jsonString)
	myurl := "https://jsonplaceholder.typicode.com/todos"
	//send POST request
	res, err := http.Post(myurl, "application/json", jsonReader)
	if err != nil {
		fmt.Println("Error Sending Request:", err)
		return
	}
	defer res.Body.Close()
	//data, _ := io.ReadAll(res.Body)
	//fmt.Println("Response :", string(data))
	fmt.Println("Response status :", res.Status)

}

func performUpdateReq() {
	todo := Todo{
		UserID:    1811,
		Title:     "Sarthak",
		Completed: false,
	}
	jsonData, err := json.Marshal(todo)
	if err != nil {
		fmt.Println("Error MArshling :", err)
		return
	}

	//Convert JSON data to string
	jsonString := string(jsonData)

	//convert string data to reader
	jsonReader := strings.NewReader(jsonString)
	const myurl = "https://jsonplaceholder.typicode.com/todos/1"

	//create PUT request
	req, err := http.NewRequest(http.MethodPut, myurl, jsonReader)
	if err != nil {
		fmt.Println("Error creating PUT Request:", err)
		return
	}
	req.Header.Set("Content-type", "application/json")

	//Send the Request
	client := http.Client{}
	res, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending request :", err)
		return
	}
	defer res.Body.Close()
	data, _ := io.ReadAll(res.Body)
	fmt.Println("Response :", string(data))
	fmt.Println("Response status :",res.Status)

}

func performDeleteReq(){
	const myurl = "https://jsonplaceholder.typicode.com/todos/3"
	//create DElETE Request 
	req, err := http.NewRequest(http.MethodDelete, myurl, nil)
	if err != nil {
		fmt.Println("Error creating Delete  Request:", err)
		return
	}
	//Send the Request
	client := http.Client{}
	res, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending request :", err)
		return
	}
	defer res.Body.Close()
	
	fmt.Println("Response status :",res.Status)

}


func main() {
	fmt.Println("Learning CRUD....")
	//performGetReq()

	//performPostReq()
	//performUpdateReq()
	performDeleteReq()

}
