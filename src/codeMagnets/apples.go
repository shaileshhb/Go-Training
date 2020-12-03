package main

import "fmt"

func main() {
	var originalCount int = 10
	var eatenCount int = 4

	fmt.Println("I started with", originalCount, "apples")
	fmt.Println("Somebody ate", eatenCount, "apples")

	eatenCount = originalCount - eatenCount

	fmt.Println("There are", eatenCount, "apples left")
}
