package main

import "fmt"

func main() {
	//name <-------> grade
	studentGrades := make(map[string]int)
	//map bna liya

	studentGrades["neetu"] = 34
	studentGrades["Alice"] = 49
	studentGrades["Bob"] = 85
	studentGrades["Charlie"] = 93

	fmt.Println("marks of Bob :", studentGrades["neetu"])
	delete(studentGrades, "Bob")
	fmt.Println("marks of Bob :", studentGrades["Bob"])

	//checking if a key exists
	grades, exists := studentGrades["neetu"]
	fmt.Println("grades of neetu: ", grades)
	fmt.Println("davis exists:", exists)

	for index, value := range studentGrades {
		fmt.Printf("key is %s and marks is %d\n", index, value)
	}

	person := map[string]int{
		"alice":   90,
		"Bob":     40,
		"Charlie": 68,
	}
	println("_____________________________")

	for index, value := range person {
		fmt.Printf("key is %s and marks is %d\n", index, value)
	}

}
