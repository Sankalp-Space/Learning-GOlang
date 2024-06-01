package main

import "fmt"

func main() {
	fmt.Println("I m Learning Array In Golang")

	/* var name [5]string
	name[0] = "Sankalp"
	name[2] = "Sarthak"

	fmt.Println("Name of person is :",name )

	var numbers= [7]int {1,2,3,4,5}
	fmt.Println("Number is:",numbers)

	fmt.Println("Lenght of Numbers array is ",len(numbers))

	fmt.Println("Value of name at second Index is :",name[2]) */

	var price[5]string 
	price[2]="Sankalp Shrivastav"
	price[0]="Sarthak Shrivastav"
	fmt.Println("Price is :",price)
	fmt.Printf("price array is %q\n", price)
}