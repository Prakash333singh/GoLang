package main

import (
	"fmt"
	"net/url"
)

///url.Parse fn is used to parse a string into a url object

func main() {
	fmt.Println("hannjjjjj kese ho!!!")
	myURl := "https://example.com:8080/path/to/resource?key1=value1&key2=value2"

	parsedURL, err := url.Parse(myURl)
	if err != nil {
		fmt.Println("error parsing Url :", err)
		return
	}

	fmt.Println("Scheme:", parsedURL.Scheme) //protocol to be used
	//http, https, ftp, mailto
	fmt.Println("host", parsedURL.Host)
	fmt.Println("Path", parsedURL.Path)         //Path /path/to/resource
	fmt.Println("RawQuery", parsedURL.RawQuery) //key1=value1&key2=value2

	/////
	//fmt.Printf("hello %T", parsedURL)
	newURL := parsedURL.String()
	fmt.Println("modified Url :", newURL)
	//fmt.Printf("%T", newURL)
}
