package main

import (
	"fmt"
	"sync"
)

func a(mg *sync.WaitGroup) {
	defer mg.Done()
	fmt.Println("process a is running....")
}

func main() {
	// wg := sync.WaitGroup()
	wg := sync.WaitGroup{}

	wg.Add(1)
	go a(&wg)

	fmt.Println("************")
	wg.wait()
	fmt.Println("&&&&&&&&&&")
}
