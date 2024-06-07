package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	fmt.Println("Learning web service ")

	res, err := http.Get("https://fakestoreapi.com/products/1")
	if err != nil {
		fmt.Println("Error in web request ", err)
		return
	}
	defer res.Body.Close()
	fmt.Printf("Type of response :%T\n", res)
	//fmt.Println("Response",res)

	// Read the response body

	data, errors := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println("Error reading response ", errors)
		return
	}
	fmt.Println("Response is:", string(data))

}
