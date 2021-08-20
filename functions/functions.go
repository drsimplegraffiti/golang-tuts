package main

import "fmt"

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

func main() {
	greet("joseph")
	greet("tye")
	sayHello("mimi")

	cycleThroughNames([]string {"bass","guitar","kong"}, sayHello)
	cycleThroughNames([]string {"bass","guitar","kong"}, greet)

}
