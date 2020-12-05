package main

import (
	"fmt"
	"runtime"
)

// variable can be declared bt cannot be assigned
// if it is not used then no error thrown
var name string

func main() {
	var rollno int
	rollno = 5
	name = "tom"
	var (
		firstName, lastName = "sam", "thomas"
	)

	fmt.Println(rollno, name)
	fmt.Println(firstName, lastName)

	// Shorthand declaration
	//type of num cannot be changed as it is set to int
	num := 10
	fmt.Printf("%T", num)

	fmt.Println(runtime.GOOS)
	fmt.Println(runtime.GOARCH)
}
