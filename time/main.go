package main

import (
	"fmt"
	"time"
)

func main() {

	currentTime := time.Now()
	fmt.Println("Current time :", currentTime)
	fmt.Printf("Type of Current Time:%T \n", currentTime)

	fomatted := currentTime.Format("2006/01/02, 03:04 PM")
	fmt.Println("Formatted time :", fomatted)

	layout_str:="02/01/2006"
	dataStr:="18/11/2002"
	formatted_time,_:=time.Parse(layout_str,dataStr)
	fmt.Println("Formatted time is:",formatted_time)

	//add one more time in Currenttime
	new_date:=currentTime.Add(24*time.Hour)
	fmt.Println("New_date time :",new_date)
	formatted_new_date:=new_date.Format("2006/01/02")
	fmt.Println("Formatted Date is :",formatted_new_date)
}