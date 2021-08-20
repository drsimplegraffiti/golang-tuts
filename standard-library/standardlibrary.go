package main

import (
	"fmt"
	"math"
	"strings"
	"sort"
)

func main() {
	greeting := "hello dev gangs"
	ages := []int{34,61,22,65,232,45,32,69}
	
	fmt.Println(math.Floor(2.5))
	fmt.Println(math.Ceil(2.5))
	fmt.Println(math.Sqrt(36))
	//@returns boolean: i.e true because it contains the string hello
	fmt.Println(strings.Contains(greeting, "hello"))
	
	//@returns boolean: i.e false because it doesnt contains the string hello
	fmt.Println(strings.Contains(greeting, "hello!"))
	//@ replace strings
	fmt.Println(strings.ReplaceAll(greeting, "hello", "hey"))

	//@ to upper case
		fmt.Println(strings.ToUpper(greeting))
	
		//@ finding index position
		fmt.Println(strings.Index(greeting, "lo"))

		//@ split strings
		fmt.Println(strings.Split(greeting, " "))

		//@ sorting integers
		sort.Ints(ages)
		fmt.Println(ages)
		
		// finding sorted position
		index := sort.SearchInts(ages, 61)
		fmt.Println(index)

		// slicing strings
		cars := []string{"volvo", "toyota", "benz", "ford"}
		sort.Strings(cars)
		fmt.Println(cars)

		//search the sliced strings...Here it sorts alphabetically and benz (0),ford (1), toyota(2), volvo(3)
		fmt.Println(sort.SearchStrings(cars, "toyota"))
}

