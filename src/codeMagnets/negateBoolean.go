package main

import "fmt"

func negate(boolValue *bool) *int {
	*boolValue = !*boolValue
	var num int = 10
	return &num
}

func main() {
	truth := true
	var num *int = negate(&truth)
	fmt.Println(truth)
	fmt.Println(*num)
	fmt.Println(num)
}
