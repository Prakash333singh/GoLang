package main

import (
	"fmt"
	"time"
)

func main() {
	currentTime := time.Now()
	formattedTime := currentTime.Format("02-01-2006,15:04:05, Monday")
	fmt.Println(currentTime, formattedTime)

	fmt.Println("start of the program")
	// The function inside defer will be executed when the surrounding
	// function (main in this case) exits
	///if there are multiple defer then they are excauted in lifo
	//order
	defer fmt.Println("this will be excauted at the end 1")
	defer fmt.Println("this will be excauted at the end 2")
	fmt.Println("middle of the program")

}
