//parallel
// chize sath sath chal rhi hai

// concurrency
// switches b/w processs bhut fastly illusion type
//tab chalta hai jab resources ka limitation rehta hai

//agar mere pass 4 core hai process mai tuo charo parrelley chal skte hai
//agar mujhe 5Th wala chalana hai tuo wo concurrent sse chelnge

package main

import (
	"fmt"
	"time"
)

func main() {
	// a() ///paraelely chate hua dikhe
	// b()
	// c()
	// d()

	go a() ///paraelely chate hua dikhe
	go b()
	go c()
	go d()

	time.Sleep(time.Second * 10)

}

func a() {
	for i := 0; i <= 3; i++ {
		fmt.Println("fun a run hua hai")
	}

}
func b() {
	for i := 0; i <= 3; i++ {
		fmt.Println("fun b run hua hai")
	}

}
func c() {
	for i := 0; i <= 2; i++ {
		fmt.Println("fun c run hua hai")
	}
}
func d() {
	for i := 0; i <= 3; i++ {
		fmt.Println("fun d run hua hai")
	}
}
