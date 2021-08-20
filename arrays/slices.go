package main

import "fmt"

func main() {
	//@ Slices
	var scores = []int{100, 50, 90,88,34,12,54,89}
	scores[1] = 44

	//append to update variable
	scores = append(scores, 46)

	fmt.Println(scores, len(scores))

	//slice ranges: in the example below it starts count from position 1 and to 2 and excludes position 3 and all
	rangeOne :=scores[1:2]
	
	//slice ranges: in the example below it starts count from position 2 and includes position the rest
	rangeTwo :=scores[2:]

	//slice ranges: in the example below it starts count from position 1 and includes to position two, leaving the rest out
	rangeThree :=scores[:3]

	fmt.Println(rangeOne)
	fmt.Println(rangeTwo)
	fmt.Println(rangeThree)

	//appending to ranges
	rangeOne =append(rangeOne, 59)
	fmt.Println(rangeOne)

}
