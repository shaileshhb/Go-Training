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

	var person2 = person{
		firstName: "Sam",
		lastName: "Thomas",
		age: 25,
		addresses: []string{
			"mumbai",
			"punjab",
		},
	}

	var person3 = person{
		firstName: "John",
		lastName: "Thomas",
		age: 28,
		addresses: []string{
			"delhi",
			"pune",
		},
	}

	fmt.Println(person1)
	modifyPerson(&person1)
	fmt.Println(person1)

	var allPersons []person
	allPersons = []person{
		person1, person2, person3,
	}

	fmt.Println(allPersons)
}

func modifyPerson(p *person){
	p.firstName = "NewName"
	p.age = 101
}

type person struct {
	firstName string
	lastName string
	age int
	addresses []string
}