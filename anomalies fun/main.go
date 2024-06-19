package main

import "fmt"

func main() {

	///have to run in routine ek sath 2 chizo ke chalane ke kaam aata hai
	func(a int) int {
		fmt.Println("anonyomous func is running...")
		return 0
	}(5)
	fmt.Println("sum =", add(5, 6, 90, 10))
}

//variadic means you can pass infinte no of varibles
func add(b ...int) int {
	sum := 0

	for _, val := range b {
		fmt.Println("val", val)
		sum += val
	}
	return sum
}
