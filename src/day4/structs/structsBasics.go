package structs

// func main() {

// 	var person1 = &person{
// 		firstName: "Sam",
// 		age:       25,
// 	}

// 	fmt.Println(*person1)

// 	var person2 = createNewPerson("Tom", 30)
// 	fmt.Println(*person2)

// }

// method for creating new structs
func CreateNewPerson(fName string, age int) *Person {
	var newPerson = &Person{
		FirstName: fName,
		Age:       age,
	}

	return newPerson
}

type Person struct {
	FirstName string
	Age       int
}
