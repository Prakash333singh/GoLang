package main

import "fmt"

//define a struct named person

type Person struct {
	FirstName string
	LastName  string
	Age       int
}

///chaining of struct
type Contact struct {
	Email string
	Phone string
}

type Address struct {
	House int
	Area  string
	state string
}

type Employee struct {
	Person_Details Person
	Person_Contact Contact
	Person_Address Address
}

func main() {
	//1st method
	// var prakash Person
	//fmt.Println("Price Person :", prakash)
	// prakash.FirstName = "prakash"
	// prakash.LastName = "singh"
	// prakash.Age = 21
	// fmt.Println("Price Person :", prakash)

	//2nd method
	// person1 := Person{
	// 	FirstName: "Aakash",
	// 	LastName:  "Singh",
	// 	Age:       26,
	// }
	//fmt.Println("Person 1 :", person1)

	//new keyword
	var person2 = new(Person)
	person2.FirstName = "neetu"
	person2.LastName = "bisht"
	person2.Age = 21

	// fmt.Println("Person 2:", person2)
	//fmt.Println("Person 2:", person2.FirstName)
	var employee1 Employee

	employee1.Person_Details = Person{
		FirstName: "PRice",
		LastName:  "Agrawal",
		Age:       25,
	}
	employee1.Person_Contact.Email = "pa78455@gmail.com"
	employee1.Person_Contact.Phone = "82900100918"

	employee1.Person_Address = Address{
		House: 12,
		Area:  "nainital",
		state: "uttrakhand",
	}

	fmt.Println(employee1)
}
