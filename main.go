package main

import "fmt"

type conctactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstName string
	lastName  string
}

func main() {
	//Ways to assign values to structs

	//#1 Declare Struct
	// alex := person{"Alex", "Anderson"}

	// #2 Declare Struct
	// alex := person{firstName: "Alex", lastName: "Anderson"}

	// #3 Declare Struct
	var alex person
	alex.firstName = "Alex"
	alex.lastName = "Anderson"

	fmt.Println(alex)
	fmt.Println(alex.firstName)
	fmt.Println(alex.lastName)
}
