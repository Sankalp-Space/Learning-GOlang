package main

import (
	"fmt"
	"time"
)

func sayHello() {
	fmt.Println("Hello,World")
	time.Sleep(3000 * time.Millisecond)
	fmt.Println("sayHello function ended successfully")
}
func sayHi() {
	fmt.Println("Hi,Sankalp :)")
	time.Sleep(2000 * time.Millisecond)
	fmt.Printf("sayHi function ended successfully")
}

func main() {
	fmt.Println("Learning Goroutines")
	go sayHello()
	go sayHi()

	time.Sleep(2500 * time.Millisecond)

}
