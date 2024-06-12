package main

import "fmt"

func divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, fmt.Errorf("denominator must not be zero")
	}
	return a / b, nil
}

func main() {
	// data := divide(67, 21)
	data, _ := divide(10.000, 0.00000)

	fmt.Println(data)
	fmt.Printf("data is %.4f\n", data)
}
