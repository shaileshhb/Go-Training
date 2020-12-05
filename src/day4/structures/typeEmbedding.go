package main

func main() {

	var emp1 = &Employee{
		EmployeeID: 1,
		Person: Person{
			FirstName: "Sam",
			Age:       26,
		},
	}

	fmt.Println(*emp1)
	fmt.Println(emp1.employeeID)
	fmt.Println(emp1.person.firstName)
}

type Employee struct {
	EmployeeID int
	Person
}

type Person struct {
	FirstName string
	Age       int
}
	FirstName string
	Age       int
}
