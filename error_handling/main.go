package main

import "fmt"

/* func divide(a, b float32) (float32, error) {
	if b == 0{
		return 0,fmt.Errorf("denominator must not be zero")
	}
	return a / b,nil */
func divide(a, b float32) (float32, string ) {
	if b == 0{
		return 0,"denominator must not be zero"
	}
	return a / b,"nil"

}
func main() {
	//fmt.Println("standard Error Handling in the function")
	//ans,err:= divide(9.2, 0)
	//if err!=nil{
	//	fmt.Println("Error Handling ")
	//}
	//fmt.Println("The division of two number is",ans)
	
	
	ans1,_:=divide(10.4,2)
	fmt.Println("The divsion is", ans1 )
}
