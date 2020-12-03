package main

import (
	"fmt"
	"reflect"
)

func main() {
	var a int = 10
	// b := "hello"
	fmt.Println("Value:", a)
	fmt.Println("Memory location:", &a)

	p := &a
	// p := &b
	fmt.Println("Value:", p)
	fmt.Println("Value of a (derefrencing of pointer):", *p)
	fmt.Println("Type:", reflect.TypeOf(p))
	fmt.Println("Memory Location", &p)

	var ptrToptr **int = &p
	fmt.Println("Value:", ptrToptr)
	fmt.Println("Value of a (derefrencing of pointer):", *ptrToptr)
	fmt.Println("Value of a (double derefrencing of pointer):", **ptrToptr)
	fmt.Println("Type:", reflect.TypeOf(ptrToptr))
	fmt.Println("Memory Location", &ptrToptr)

	fmt.Println("\n======================================")
	var stringValue = "Hello"
	floatValue := createPointer(&stringValue)
	fmt.Println("Float value: ", *floatValue)
}

func createPointer(stringValue *string) *float64 {
	var floatValue = 20.3
	*stringValue += " world"
	fmt.Println(*stringValue)
	return &floatValue
}
