package main

import (
	"bufio"
	"fmt"
	"os"
)

func add(a, b int) (result int) {
	result = a + b
	return result
}

func main() {
	// age := 25
	// name := "alice"
	// height := 5.8992728

	// fmt.Println("age:", age, "height:", height, "name:", name)
	// fmt.Println("hmmmm ji kesi ho")

	// fmt.Printf("age is %d\n", age)
	// fmt.Printf("height is %.4f\n", height)
	// fmt.Printf("Type of age is %T\n", age)
	// fmt.Printf("name is %s\n", name)

	// fmt.Println("hey, what's your name?")
	// var name string
	// fmt.Scan(&name)

	// fmt.Println("hello", name)

	//if i want to long input string which has space

	fmt.Println("hey, what's your name?")
	reader := bufio.NewReader(os.Stdin)
	name, _ := reader.ReadString('\n')
	fmt.Println("hello .Mr ", name)

	var ans = add(9, 10)
	fmt.Println(ans)
}
