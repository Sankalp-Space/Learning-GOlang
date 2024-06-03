package main

import "fmt"

func main() {
	for i := 0; i < 1; i++ {
		fmt.Println("Numbers is :",i)
	}
	//infinite loop with break statement 
	counter :=0
	for{
		fmt.Println("Infinite loop")
		counter++
		if counter==2{
			break
		}
	}
	numbers:= []int{21,82,34,46,54}
	for index,value:=range numbers{
		fmt.Printf("Index of number is %d, and value is %d \n",index,value)
	}
	data := "Hello,World"
	for index,char:=range data{
		fmt.Printf("Index of number is %d, and value is %c \n",index,char)
	}

}