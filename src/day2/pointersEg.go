package main

import (
	"fmt"
	"reflect"
)

func main() {
	var a int = 10
	fmt.Println("Value:", a)
	fmt.Println("Memory location:", &a)

	p := &a
	fmt.Println("Value:", p)
	fmt.Println("Value of a (derefrencing of pointer):", *p)
	fmt.Println("Type:", reflect.TypeOf(p))
	fmt.Println("Memory Location", &p)
}
