package main

//by default nil pointers are set to valid memory address

import "fmt"

func modifyValueByReference(num *int) {
	*num = *num * 2
}

func main() {
	// var ptr *int

	// if ptr == nil {
	// 	fmt.Println("pointer to null")
	// }

	value := 100
	modifyValueByReference(&value)
	fmt.Println("Modified value", value)

}
