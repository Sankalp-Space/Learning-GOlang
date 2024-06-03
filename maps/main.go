package main

import "fmt"

func main() {
	//name with grade
	stdgrades := make(map[string]int)
	stdgrades["Sankalp"] = 99
	stdgrades["Sarthak"] = 98
	stdgrades["Prince"] = 100
	stdgrades["Bittu"] = 33

	fmt.Println("Marks of Bittu :",stdgrades["Bittu"])

	stdgrades["Bittu"]= 96
	fmt.Println("New marks of Bittu is:",stdgrades["Bittu"])

	delete(stdgrades,"Bittu")
	fmt.Println("New marks of Bittu is:",stdgrades["Bittu"])
	//checking if a key exists or not

	grade,exists:=stdgrades["David"]
	fmt.Println("Grades of Aayush:",grade)
	fmt.Println("Aayush exist ",exists)
	
	grades,exist:=stdgrades["Sarthak"]
	fmt.Println("Grades of Sarthak:",grades)
	fmt.Println("Sarthak exist ",exist)
	for index,value:=range stdgrades{
		fmt.Printf("Key is %s and marks is %d\n",index ,value)
	}


}