package main

import "fmt"


func main() {
	myAge := 34

	//Using conditional codes
	if myAge < 34 {
fmt.Println("age is less than 34")
	} else if myAge < 46 {
		fmt.Println("age is less than 46")
	} else {
		fmt.Println("age is not less than 46")
	}

	names := []string{"bimbo", "tye", "rick", "james"}

	for index, value := range names{
		if index == 1 {
			fmt.Println("continuing at pos",index)
			continue
		}
		if index > 2{
			fmt.Println("breaking at pos", index)
			break
		}
		fmt.Printf("the value at pos %v is %v \n", index, value)
	}
}
