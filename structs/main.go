package main

import "fmt"

type person struct {
	firstName string
	lastName  string
	contactInfo
}

type contactInfo struct {
	email   string
	zipCode int
}

func main() {
	// // Only person struct demmo
	// alex := person{"Alex", "Anderson"}
	// alex := person{firstName: "Alex",
	// 	lastName: "Anderson"}
	// var alex person
	// alex.firstName = "Alex"
	// alex.lastName = "Anderson"

	akhil := person{
		firstName: "Akhil",
		lastName:  "Yethi",
		contactInfo: contactInfo{
			email:   "test@gmail.com",
			zipCode: 94000,
		},
	}

	akhil.print()
	akhil.updateName("Akku")
	akhil.print()
}

func (p person) print() {
	fmt.Printf("%+v\n", p)
}

func (p person) updateName(newFirstName string) {
	p.firstName = newFirstName
}
