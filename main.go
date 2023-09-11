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

	jim.updateName("Jeong")
	jim.print() // Jim
}

// The original Instance have not changed by this.
// So, You should use a pointer for it to work as expected
//func (p person) updateName(newFirstName string) person {
//	p.firstName = newFirstName
//	return p // pass by value
//}

// Go will allow us to either call this function with a pointer or with the root type
func (p *person) updateName(newFirstName string) {
	p.firstName = newFirstName
}

func (p *person) print() {
	fmt.Printf("%+v", p)
}
