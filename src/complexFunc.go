package main

import (
	"fmt"
)

func main() {
	fmt.Println(test) // Print address
	// fmt.Println(test(10))

	str := abc(addOne)
	fmt.Println(str())

}

func addOne(num int) int {
	num = num + 1
	return num
}

func greetings() string {
	return "Hey there!"
}

func test(num int) {

}

func abc(func1 func(num1 int) int) func() string {
	fmt.Println(func1(10))
	return greetings
}
