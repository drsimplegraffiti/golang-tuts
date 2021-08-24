package main

import "fmt"

func main() {
	menu := map[string]float64{
		"stew": 4.88,
		"mango": 5.87,
		"cake": 6.43,
	}
	fmt.Println(menu)
	fmt.Println(menu["cake"])

	//looping through maps
	for k, v := range menu {
		fmt.Println(k, "-", v)
	}

	//ints as key type
	phonebook := map[int]string{
		944343455: "james",
		3434343455: "gawk",
		3445893455: "lord",
	}
	fmt.Println(phonebook)
	fmt.Println(phonebook[944343455])
	
	
	//update items
	phonebook[3434343455] = "jordan"
	fmt.Println(phonebook)
	
	
	phonebook[944343455] = "tim"
	fmt.Println(phonebook)
}
