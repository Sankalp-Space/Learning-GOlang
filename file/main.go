package main

import (
	"fmt"
	"os"
)

func main() {
	/* file,err:=os.Create("Example.txt")
	if err!=nil{
		fmt.Println("Error while Creating a file",err)
		return
	}
	defer file.Close()

	content := "How r u ? "
	byte , err1 := io.WriteString(file,content+"\n")
	fmt.Println("No of bytes ",byte)
	if err1!=nil{
		fmt.Println("Content was not written in file due to error",err1)
	}


	fmt.Println("Successfully created file") */

	/* file,err:=os.Open("example.txt")
	if err!=nil{
		fmt.Println("Error while Opening the file",err)
		return
	}
	defer file.Close()

	//creating a buffer to read the file content

	buffer:=make([]byte,1024)

	//Read the file content into the buffer
	for{
		n,err:=file.Read(buffer)
		if err == io.EOF{
			break

		}
		if err!=nil{
			fmt.Println("Error while reading file",err)
			return
		}
		// process to read the  content
		fmt.Println(string(buffer[:n]))
	} */

	content, err := os.ReadFile("Example.txt")
	if err != nil {
		fmt.Println("Error while reading file", err)
		return
	}
	fmt.Println(string(content))

}
