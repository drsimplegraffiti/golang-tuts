package main

import (
	"fmt"
)

var age = 456

func main() {
	sayHello("joseph")

	for _, v := range points {
		fmt.Println(v)
	}
	showAge()
}


//running two files go run main.go greetings.go