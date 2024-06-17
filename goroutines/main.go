package main

// Goroutines are a powerful feature in Go that make it easy to run
// functions concurrently, significantly improving the performance
// and responsiveness of your programs.
//  In the example above, using goroutines allows a web server to
// handle multiple requests simultaneously

import (
	"fmt"
	"net/http"
	"time"
)

func handler(w http.ResponseWriter, r *http.Request) {
	start := time.Now()

	// Simulate a time-consuming task
	time.Sleep(2 * time.Second)

	// Use the ResponseWriter to send a response back to the client
	fmt.Fprintf(w, "Hello, World! The current time is %s\n", time.Now().Format(time.RFC1123))

	// Log the time taken to handle the request
	fmt.Printf("Handled request in %s\n", time.Since(start))
}

// func sayHello() {
// 	time.Sleep(10000 * time.Millisecond)
// 	fmt.Println("haanjji")
// }

// func sayHi() {
// 	time.Sleep(1000 * time.Millisecond)
// 	fmt.Println("gudmorng")
// }

func main() {
	//start a new goroutines
	// go sayHello()
	// sayHi()

	//wait for a moment to allow the goroutine to finish
	//time.Sleep(11000 * time.Millisecond)

	// Register the handler function for the "/" route
	http.HandleFunc("/", handler)

	// Start the web server on port 8080
	fmt.Println("Starting server on :8080")
	// The ListenAndServe function starts a new goroutine for each incoming request
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Printf("Error starting server: %s\n", err)
	}
}
