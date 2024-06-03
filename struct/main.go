package main

import "fmt"

type Person struct {
	FirstName string
	LastName  string
	Age       int
}
type contact struct{
	Email string 
	Phone int

}
type Address struct{
	House int 
	Area string 
	State string 
}

type Employee struct{
	Person_Details Person
	Person_Contact contact
	Person_Address Address

}

func main() {
	var sankalp Person
	//fmt.Println("Prince  person ",sankalp)
	sankalp.FirstName="Sankalp"
	sankalp.LastName="Shrivastav"
	sankalp.Age=21
	fmt.Println("sankalp  details",sankalp)

	var sarthak Person
	sarthak.FirstName="Sarthak"
	sarthak.LastName="Shrivastav"
	sarthak.Age=19
	fmt.Println("Sarthak detail",sarthak)

	Prince :=Person{
		FirstName: "Prince",
		LastName: "Kulshertha",
		Age: 18,
	}
	fmt.Println("Prince Details:",Prince)

	//new keyword
	var simran=new(Person)
	simran.FirstName="Simran"
	simran.LastName="SIMRAN"
	simran.Age=20

	//fmt.Println("Simran Lastname was:",simran.LastName)
	//fmt.Println("Sarthak Age is :",sarthak.Age)

	var employee1 Employee
	employee1.Person_Details=Person{
		FirstName: "Abc",
		LastName: "Ced",
		Age: 26,
	}
	employee1.Person_Contact.Email="contact@gamil.com"
	employee1.Person_Contact.Phone=9876543210

	employee1.Person_Address= Address{
		House: 101,
		Area: "Posh Area",
		State: "Maharastra",
	}
	fmt.Println("Employee 1",employee1)
	fmt.Println("Employee 1 contacts",employee1.Person_Contact)

}