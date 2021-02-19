package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstName string
	lastName  string
	contact   contactInfo
}

func main() {
	//Ways to assign values to structs

	//#1 Declare Struct
	// alex := person{"Alex", "Anderson"}

	// #2 Declare Struct
	// alex := person{firstName: "Alex", lastName: "Anderson"}

	// #3 Declare Struct
	// var alex person
	// alex.firstName = "Alex"
	// alex.lastName = "Anderson"

	jim := person{
		firstName: "Jim",
		lastName:  "Party",
		contact: contactInfo{
			email:   "jim@gmail.com",
			zipCode: 94000,
		},
	}

	fmt.Printf("%+v", jim)

	fmt.Println(jim)
	fmt.Println(jim.firstName)
	fmt.Println(jim.lastName)
}
