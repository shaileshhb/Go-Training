package main

import "fmt"

func main() {

	number := 10
	fmt.Println("Address in main: ", &number)
	passByValue(number)
	fmt.Println("Value after passByValue", number)

	passByReference(&number)
	fmt.Println("Value after passByReference", number)

}

// Pass By Value
func passByValue(num int) {
	num += 2
	fmt.Println("Value in value: ", num)
	fmt.Println("Address in value: ", &num)
}

// Pass By Reference
func passByReference(num *int) {
	*num += 2
	fmt.Println("Value in reference: ", *num)
	fmt.Println("Address in reference: ", num)
}
