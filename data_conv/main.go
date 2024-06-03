package main

import (
	"fmt"
	"strconv"
)

func main() {
	var num int = 42
	fmt.Println("Number is ",num)
	fmt.Printf("Type of number : %T \n",num)

	var data float64 = float64(num)
	data = data +1.23
	fmt.Println("Data is :",data)
	fmt.Printf("Type of data : %T\n",data)

	num =122
	str:=strconv.Itoa(num)
	fmt.Println("Str is :",str)
	fmt.Printf("Type of str is  %T \n",str)

	number_str:="1234"
	number_int,_:=strconv.Atoi(number_str)
	number_int=number_int+1234
	fmt.Println("Number interger is :",number_int)
	fmt.Printf("Type of Number integer is %T\n",number_int)

	num_str:="3.14"
	num_flt,_:=strconv.ParseFloat(num_str,64)
	fmt.Println("Number float is :",num_flt)
	fmt.Printf("Type of Number integer is %T",num_flt)

}