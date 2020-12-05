package main

import (
	"fmt"
)

func main() {
	displayName()
	userNumber := addOne(10)
	fmt.Println(userNumber)
	fmt.Println("Inside Main")
}

func displayName() {
	fmt.Println("Inside displayName")
}

func addOne(number int) int {
	number += 1
	return number
}
