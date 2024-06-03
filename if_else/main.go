package main

import "fmt"

func main() {
	x := 0
	if x == 0 {
		fmt.Println("Its a Zero")
	}else{
		fmt.Println("The number is non zero")


	}

	z:=10
	if z>=10{
		fmt.Println("z is greater than 10 or equal to 10")

	}else if z>5 {
		fmt.Println("z is greater than 5 but smaller than 10")
		
	}else{
		fmt.Println("z is smaller than 5")
	}

	y:=2
	if x>5 && y>5{
		fmt.Println("both x and y is greater than 5")
	}else{
		fmt.Println("Both x and y is less than 5")
	}
}