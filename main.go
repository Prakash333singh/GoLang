package main

import "mylearning/mypackage"

func main() {
	// var name string = "adhikari"
	// fmt.Println(name)
	// fmt.Println("Pagal hai kya ")
	// myutil.PrintMessage("chal na ")

	// var person = 23
	// person = 20
	// fmt.Println(person)

	// const person = 23
	// person = 20 can't be redeclare
	// fmt.Println(person)

	// person := 123
	// person := "hello"
	// fmt.Println(person)

	mypackage.PublicFunction()
	mypackage.PrivateFunction()
	//privateFunction()

}

//mylearning/
// ├── go.mod
// ├── main.go
// └── myutil/
//     └── myutil.go
