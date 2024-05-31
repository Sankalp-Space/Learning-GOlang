package main

import "fmt"

func simpleFucntion(){
	fmt.Println("Simple Function")
}
func add(a,b int) (result int){
	result= a+b
	return 
}
func multiply(a,b int )(result int){
	result= a*b
	return 
}

func main() {
	fmt.Println("we are learing Funtion in Golang")
	simpleFucntion()

	ans:=add(3,5)
	ans1:=multiply(4,5)
	fmt.Println("add of two numer is :",ans)
	fmt.Println("multiplication of two numer is :",ans1)


}