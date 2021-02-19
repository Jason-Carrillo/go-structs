package main

import "fmt"

type person struct {
	firstName string
	lastName  string
}

func main() {
	//Ways to assign values to structs
	// alex := person{"Alex", "Anderson"}

	alex := person{firstName: "Alex", lastName: "Anderson"}

	fmt.Println(alex)
	fmt.Println(alex.firstName)
	fmt.Println(alex.lastName)
}
