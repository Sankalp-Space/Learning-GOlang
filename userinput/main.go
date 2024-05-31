package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println("Hey , What's your name ?")
	//var name string;

	//fmt.Scan(&name )
	//fmt.Println("Hello ",name)

	reader := bufio.NewReader(os.Stdin)
	name, _ := reader.ReadString('\n')
	fmt.Println("Hello, Mr.", name, "how are you")
	read := bufio.NewReader(os.Stdin)
	gret,_:=read.ReadString('\n')
	fmt.Print("okkk",gret,"have a great day")

}
