package main

import (
	"fmt"
	"net/url"
)

func main() {
	fmt.Println("Learing URL")
	myurl:="https://study.com/academy/lesson/what-is-a-url-definition-examples-quiz.html"
	fmt.Printf("Type of URL: %T \n",myurl)

	parseURL,err:=url.Parse(myurl)
	if err!=nil{
		fmt.Println("Can't parse URL",err)
		return
	}
	fmt.Printf("Type of URL:%T\n",parseURL)
	fmt.Println("Scheme :",parseURL.Scheme)
	fmt.Println("Host :",parseURL.Host)
	fmt.Println("Path :",parseURL.Path)
	fmt.Println("RawQuery :",parseURL.RawQuery)


	//Modifying URL components
	parseURL.Path="/newpath"
	parseURL.RawQuery="username=sankalp00"

	//constructing a URL string from URL objects 
	newUrl:=parseURL.String()
	fmt.Println("new URL :",newUrl)
}