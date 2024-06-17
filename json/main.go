package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Name    string `json:"name"`
	Age     int    `json:"age"`
	IsAdult bool   `json:"is_adult"`
}

func main() {
	// Encoding (Marshalling)
	person := Person{Name: "Alice", Age: 30, IsAdult: true}
	jsonData, err := json.Marshal(person)
	if err != nil {
		fmt.Println("Error marshalling JSON:", err)
		return
	}
	fmt.Println("JSON data:", string(jsonData))

	// Decoding (Unmarshalling)
	var newPerson Person
	err = json.Unmarshal(jsonData, &newPerson)
	if err != nil {
		fmt.Println("Error unmarshalling JSON:", err)
		return
	}
	fmt.Println("Decoded Person:", newPerson)
}
