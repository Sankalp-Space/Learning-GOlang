package main

import "fmt"

func modifyvalue(num *int){
	*num=*num +20

}

func main() {
	num:="sankalp"
	/* var num int
	num = 2

	var ptr *int
	ptr = &num */
	ptr:=&num



	//fmt.Println("Num has value :",num)
	fmt.Println("Pointer contains :",ptr)
	fmt.Println("The value at ptr",*ptr)

	var pointer *int//default pointer value is nil
	if pointer == nil{
		fmt.Println("The pointer is not assigned")
	} 
	value:=10
	modifyvalue(&value)
	fmt.Println("Value Contain :",value)

}