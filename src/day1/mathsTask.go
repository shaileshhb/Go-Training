package main

import "fmt"

func main() {

	number := mathOperations(10, 5, add)
	fmt.Println(number)

}

func add(num1, num2 int) int {
	num := num1 + num2
	return num
}

func subtract(num1, num2 int) int {
	num := num1 - num2
	return num
}

func multiply(num1, num2 int) int {
	num := num1 * num2
	return num
}

func mathOperations(num1, num2 int, funcOperation func(num1, num2 int) int) int {

	num := funcOperation(num1, num2)
	return num
}
