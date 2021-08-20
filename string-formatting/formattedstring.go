package main

import "fmt"

func main() {
	//variables
	name := "jay"
	age := 45
	nameOne := "john"
	ageOne  := 55
	decimalNumber := 92.4


	//@ new line character \n


	//formatted string: Embed using a format specifier (%_)
	// The order matters
	fmt.Printf("my name is %v and my age is %v \n", name,age)


	//%q place quotes on outputs
	fmt.Printf("my name is %q and my age is %q \n", nameOne,ageOne)
	fmt.Printf("my name is %q and my age is %v \n", nameOne,ageOne)
	
	// Variable type
	fmt.Printf("my name is of type %T \n", name)
	//@floats
	fmt.Printf("my name is of type %f \n", decimalNumber)
	//@ limit number of decimal points: 0.1 means 1 decimal point, 0.2 means 2 decimal point etc
	fmt.Printf("my name is of type %0.1f \n", decimalNumber)
	//@ pass in variable directly
	fmt.Printf("my age is of type %f \n", 34.56)


	//Sprintf (saves formatted strings)
	var store = fmt.Sprintf("my name is %v and my age is %v \n", name,age)
	fmt.Println("the stored data is :", store)
}
