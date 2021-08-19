package main

import (
	"fmt"
)

func main() {
	//@ loop/cycle through strings using the range loop
	names := []string{"boy", "girl", "pet", "alien"}
	for index, value := range names{
		fmt.Printf("the position value at index %v is %v \n", index, value)
	}

	// if you dont want to use the value or indx, replace with _
	for _, value := range names{
		fmt.Printf("the position value is %v \n",value)
	}
}
