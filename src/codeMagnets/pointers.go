package main

import (
	"fmt"
	"reflect"
)

func main() {
	var num int = 10
	var numPointer *int
	numPointer = &num
	fmt.Println(numPointer)
	fmt.Println(reflect.TypeOf(&num))
	fmt.Println(*numPointer)
	*numPointer = 15
	fmt.Println(*numPointer, num)

	var floatNum float64 = 10.4
	floatNumPointer := &floatNum
	fmt.Println(floatNumPointer)
	fmt.Println(reflect.TypeOf(&floatNum))
	fmt.Println(*floatNumPointer)

	// var myInt int
	// var myIntPointer *int = &myInt

	// *myIntPointer = 42
	// fmt.Println(myInt)
}
