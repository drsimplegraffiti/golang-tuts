package main

import (
	"fmt"
	"math"
)

func greet(n string){
	fmt.Printf("Good evening %v \n", n)
}
func sayHello(n string){
	fmt.Printf("I am saying hello to : %v \n", n)
}

// @desc Taking multiple arguments
func cycleThroughNames(n []string, f func(string)){
	for _, v := range n {
		f(v)
	}
}

//returning a value
func circleArea(r float64) float64{
	return math.Pi * r * r
}

func main() {


	cycleThroughNames([]string {"bass","guitar","kong"}, sayHello)
	cycleThroughNames([]string {"bass","guitar","kong"}, greet)

	a1 := circleArea(10.3)
	a2 := circleArea(20)
	fmt.Println(a1,a2)
	fmt.Printf("circle 1 is %0.3f and circle 2 is %0.3f",a1,a2)
}
