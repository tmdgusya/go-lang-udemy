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
	jim := person{
		firstName: "Jim",
		lastName:  "Party",
		contact: contactInfo{
			email:   "jim@party.com",
			zipCode: 50500,
		},
	}

	jim.print()
}

func (p person) print() {
	fmt.Printf("%+v", p)
}
