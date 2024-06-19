package main

import (
	"fmt"
)

type User struct {
	Name    string
	Address string
	Age     int
}

func main() {
	fmt.Println("helllllll0")
	var vishal User
	vishal.addUser(8)
	//addUser(4)
}

//kisi bhi function ke aage struct attach karne se func method mai
//convert ho jata hai

//security increase karne ke liye

func (user User) addUser(a int) {

}
