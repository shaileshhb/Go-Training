package main

import "fmt"

func main() {

	var person1 = person{
		firstName: "Tom",
		lastName: "Thomas",
		age: 30,
		addresses: []string{
			"mumbai",
			"pune",
		},
	}

	fmt.Println(person1)
	fmt.Println(person1.getFullName())
}

// this becomes a method inside person
func (p *person) getFullName() string {
	p.firstName = "Sam"
	return p.firstName + " " + p.lastName
}

type person struct {
	firstName string
	lastName string
	age int
	addresses []string
}