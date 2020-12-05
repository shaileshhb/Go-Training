package main

import (
	"fmt"
	"day4/structs"
)

func main() {

	var person1 = structs.CreateNewPerson("Sam", 25)

	fmt.Println(*person1)

	var person2 = &structs.Person{
		FirstName : "Tom",
		Age: 40,
	}

	fmt.Println(*person2)
}