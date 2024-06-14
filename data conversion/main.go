package main

import (
	"fmt"
	"strings"
)

func main() {
	// var integerNum int = 82
	// var floatNum float64 = float64(integerNum)
	// fmt.Println(floatNum)
	// var number int = 42
	// str := strconv.Itoa(number) // Integer to string
	// strNum := "123"
	// num, err := strconv.Atoi((strNum)) //string to integer

	// fmt.Println(str, num, err)
	// fmt.Printf("%T %T %T", str, num, err)
	// str := "3.14"
	// num, err := strconv.ParseFloat(str, 64)
	// if err == nil {
	// 	fmt.Println(num)
	// }

	///strings
	str1 := "apple,orange,banana,banana,banana"
	str2 := "world"
	parts := strings.Split(str1, ",")

	count := strings.Count(str1, "banana")
	result := strings.Join([]string{str1, str2}, " ")
	fmt.Println(parts, count, result)

}
