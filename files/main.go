package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	// file, error := os.Create("example.txt")
	// if error != nil {
	// 	fmt.Println("Error creating file :", error)
	// 	return
	// }
	// fmt.Println(file)
	// defer file.Close()
	// // Initial content to be added to the file
	// initialContent := "Hello, Neetu kesi hai."

	// _, error = io.WriteString(file, initialContent+"\n")

	// if error != nil {
	// 	fmt.Println("Error writing to file :", error)
	// 	return
	// }

	// fmt.Println("file created with initial content.")

	// Open the file
	// file, err := os.Open("example.txt")
	// if err != nil {
	// 	fmt.Println("Error while opening file:", err)
	// 	return
	// }

	// defer file.Close()

	// ///craete a buffer to read the file content
	// buffer := make([]byte, 1024)

	// //read the file content into the buffer
	// for {
	// 	n, err := file.Read(buffer)
	// 	if err == io.EOF {
	// 		break
	// 	}
	// 	if err != nil {
	// 		fmt.Println("Error while reading file", err)
	// 		return
	// 	}

	// 	//process the read content

	// 	fmt.Println(string(buffer[:n]))
	// }

	//read the entire file into a byte slice
	content, err := ioutil.ReadFile("example.txt")
	if err != nil {
		fmt.Println("Error while reading file", err)
		return
	}

	fmt.Println(string(content))

}
