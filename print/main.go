package main

import (
	"fmt"
)

func main() {
	age := 19
	name := "Sarthak Shrivastav"
	height := 6.144464
	fmt.Println("age : ", age, "height :", height, "name : ", name)
	//fmt.Println("new line ")


	//fmt.Printf("age is %d\n",age)
	//fmt.Printf("Height is %.2f\n",height)
	//fmt.Printf("Type of name is %T\n",name )
	//fmt.Printf("Type of height is %T\n", height)

	fmt.Printf("Name is %s\n",name)
	fmt.Printf(" Name : %s, Age : %d, Height : %.2f\n", name,age,height)

}
