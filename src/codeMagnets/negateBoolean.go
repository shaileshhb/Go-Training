package main

import "fmt"

func negate(boolValue *bool) {
	*boolValue = !*boolValue
}

func main() {
	truth := true
	negate(&truth)
	fmt.Println(truth)
}
