package main

import "fmt"

func main() {
	// Print doesn't add a new line
	fmt.Print("hello, ")
	fmt.Print("charlie")
	fmt.Print("doe")
	
	// Print @add a new line using the \n "escape"
	fmt.Print("hello, ")
	fmt.Print("charlie \n")
	fmt.Print("a new line \n")

	//PrintLine function
	fmt.Println("this is a line")
	fmt.Println("this is a new line")
}
