package main

import "fmt"

func main() {
//string variable
	var myName string = "john doe"
	var secondName = "Joseph"
	var thirdName  string
	
	fmt.Println(myName, secondName, thirdName)
	
	// @ changing values in variable
	myName = "miracle"
	secondName = "michael"
	
	fmt.Println(myName, secondName, thirdName)
	
/* 	initializing variable without the var keyword (shorthand)
	Should be used the first time when declaring variable
	Cannot be used outside of a function
*/	

 fourthName := "james"
	fmt.Println(fourthName)
	
	//Intergers in Golang
	var numberOne = 4
	var numberTwo = 100
	numberThree := 89

	fmt.Println(numberOne, numberTwo, numberThree)
	
	//Bits and Memory
	//Int refer to the int docs
	var numOne int8 =  25 
	var numTwo int8 =  -128 
	
	//Unsigned Int	means we can have a negative number
	var numThree uint8 =  255 
	fmt.Println(numOne, numTwo, numThree)
	
	//floats: bit size is very important. float32 or float64
	//float64 has higher precision than float32
	var invoiceOne  float32 = 23.9
	var invoiceTwo  float64 = 7673463746.9
		invoiceThree := 3.5 // automatically float64
	fmt.Println(invoiceOne, invoiceTwo, invoiceThree)
}
