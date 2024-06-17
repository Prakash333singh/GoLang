package main

import (
	"fmt"
	"strconv"
)

func main() {
	var integerNum int = 82
	var floatNum float64 = float64(integerNum)
	fmt.Printf("type of data is %T", floatNum)

	// var number int = 42
	// str := strconv.Itoa(number) // Integer to string
	// strNum := "123"
	// num, err := strconv.Atoi((strNum)) //string to integer

	// fmt.Println(str, num, err)
	// fmt.Printf("%T %T %T", str, num, err)

	number_string := "1234"
	number_int, _ := strconv.Atoi(number_string)
	number_int = number_int + 263689
	fmt.Println("number_int is", number_int)
	fmt.Printf("Type of number_int is %T\n", number_int)

	num_string := "3.14"
	number_float, _ := strconv.ParseFloat(num_string, 64)
	fmt.Println("number_float is ", number_float)
	fmt.Printf("Type of number_float is %T\n", number_float)

	// str := "3.14"
	// num, err := strconv.ParseFloat(str, 64)
	// if err == nil {
	// 	fmt.Println(num)
	// }

	///strings
	// str1 := "apple,orange,banana,banana,banana"
	// str2 := "world"
	// parts := strings.Split(str1, ",")

	// count := strings.Count(str1, "banana")
	// result := strings.Join([]string{str1, str2}, " ")
	// fmt.Println(parts, count, result)

}
