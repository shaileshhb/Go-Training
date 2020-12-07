package main

import "fmt"

func main() {

	var i interface{} = "hello"
	str, ok := i.(string)
	fmt.Println(str, ok)

	checkType(21)
}

func checkType(anyType interface{}) {

	switch anyType.(type) {
	case int:
		fmt.Println("Arg is string")
	case string:
		fmt.Print("string")
	case float64:
		fmt.Println("Float64")
	default:
		fmt.Println("Other type")
	}
}
