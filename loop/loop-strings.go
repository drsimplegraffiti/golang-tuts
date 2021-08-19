package main

import (
	"fmt"
)

func main() {
	//@ loop/cycle through strings
	names := []string{"boy", "girl", "pet", "alien"}
	for i := 0; i <len(names); i++ {

		fmt.Println(names[i])
	}
}
