package main

import "fmt"

func main() {
	//arrays/array variable
	//array has to be fixed length and can't be redefined
	var ages [3]int = [3]int{34,56,78} //i.e an array of three items and data type of integer

	//@ array shorthand
	var ageOne = [3]int{45,34,54}

	//@ array names
	names := [5]string{"joe","tom","jay","james","tay"}

	//@ output: age
	fmt.Println(ages,len(ages))
	fmt.Println(ageOne,len(ageOne))
	
	//@ output: name
	fmt.Println(names,len(names))
	
	//@desc Changing arrays  value
	names[2] ="rock"
	
	//@output changed array value
	fmt.Println(names,len(names))

}
