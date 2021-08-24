package main

import (
	"fmt"
	"strings"
)


func getInitials(n string)(string, string){
	// to capitalize
	s := strings.ToUpper(n)
	// split strings into an array
	names := strings.Split(s, " ")
	// create slice for the initials
	var initials []string
	for _, v := range names {
		initials = append (initials, v[:1])
	}

	if len(initials) > 1 {
		return initials[0], initials[1]
	}
	return initials[0], "_"
}


func main() {
	fn1, sn1 := getInitials("john doe")
	fmt.Println(fn1, sn1)
	fn2, sn2 := getInitials("mimi smith")
	fmt.Println(fn2, sn2)
	fn3, sn3 := getInitials("snow")
	fmt.Println(fn3, sn3)
}
