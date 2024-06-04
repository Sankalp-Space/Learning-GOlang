package main

import "fmt"

func add(a,b int)int{
	return a+b

}
func main() {
	fmt.Println("Starting of the Program ")
	data:=add(5,6)
	defer fmt.Println("The addition of a and b is :",data)
	defer fmt.Println("Middle of the Program ")//agr do se jyada defer hote he to wo LIFO follow krte he
	fmt.Println("End of the Program")
}
