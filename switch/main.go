package main

import "fmt"

func main() {
	 day := 30
//switch statement with value
	switch day {
	case 1,30,20:
		fmt.Println("Monday")
	case 2:
		fmt.Println("Tuesday")
	case 3:
		fmt.Println("Wednesday")

	default:
		fmt.Println("Unknown Day")

	} 
//switch statement with multiple value 
	month:="April"
	switch month{
	case "January","February","March":
		fmt.Println("Winter")
	case "April","May","June":
		fmt.Println("Spring")
	default:
		fmt.Println("Other Season")
	}
//switch statement with expression 
	temp :=25
	switch {
	case temp<0:
		fmt.Printf("Freezing")
	case temp>=0 && temp<10:
		fmt.Println("Cold")

	case temp>=20 && temp<30:
		fmt.Println("Warm")

	default:
		fmt.Println("Hot")


	}

}
