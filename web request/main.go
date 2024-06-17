package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	fmt.Println("learing web service")
	res, err := http.Get("https://jsonplaceholder.typicode.com/todos/1")
	if err != nil {
		fmt.Println("Error getting Get res", err)
	}

	defer res.Body.Close()

	//fmt.Printf("Type of response : %T", res)
	//fmt.Printf(" response", res)

	//read the response body
	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println("error reading response", err)
		return
	}
	fmt.Println("response :", string(data))
}
