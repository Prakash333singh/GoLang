package mypackage

import "fmt"

//this function can be used outside this folder
func PublicFunction() {
	fmt.Println("this is a public function")
}

//if i write the function name in small letter it is not accisible
///outside the dir

func PrivateFunction() {
	fmt.Println("hello this is private function")

}
