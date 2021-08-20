package main

import "fmt"


func main() {
	myAge := 34

	fmt.Println(myAge <= 45)
	fmt.Println(myAge >= 45)
	fmt.Println(myAge == 34)
	fmt.Println(myAge != 59)

	//Using conditional codes
	if myAge < 34 {
fmt.Println("age is less than 34")
	} else if myAge < 46 {
		fmt.Println("age is less than 46")
	} else {
		fmt.Println("age is not less than 46")
	}
}
