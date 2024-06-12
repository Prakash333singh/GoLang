package main

import (
	"fmt"
)

func main() {
	// numbers := []int{1, 2, 3, 4, 5}
	// numbers = append(numbers, 3, 10, 12, 13, 14, 15, 16, 17, 18, 19)
	// fmt.Println("Number: ", numbers)
	// fmt.Printf("Number has a data type :%T\n", numbers)
	// fmt.Println("Length: ", len(numbers))

	///initailze with 0
	// name := []string{}
	// fmt.Println("name: ", name)

	// numbers := make([]int, 3, 5)
	///size 3
	//capacity

	stock := make([]string, 0)

	fmt.Println("slice", stock)
	fmt.Println("Length", len(stock))
	fmt.Println("capacity", cap(stock))

	// numbers = append(numbers, 4)
	// numbers = append(numbers, 98)
	// numbers = append(numbers, 6)

	// fmt.Println("slice: ", numbers)
	// fmt.Println("Length", len(numbers))
	// fmt.Println("capacity ", cap(numbers))
}
