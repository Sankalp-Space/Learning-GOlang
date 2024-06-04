package main

import (
	"fmt"
	"strings"
)

func main() {
	data := "apple.orange.banana.mango"
	parts := strings.Split(data, ".")
	fmt.Println(parts)

	str:="one two three four  two five "
	count:=strings.Count(str,"two")
	fmt.Println("count of two is :",count)

	str ="    Hello, Go!  "
	trimmed:= strings.TrimSpace(str)
	fmt.Println("Trimmed:",trimmed)

	str1:="Sankalp"
	str2:="Shrivastav"
	result:=strings.Join([]string{str1,str2}," ....")
	fmt.Println("Result :",result)

}