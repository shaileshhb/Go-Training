package main

import "fmt"

const myConstant = 20

const (
	num1 = iota
	num2
	num3
	num4
	num5
)

func main() {

	fmt.Println("myConstant:", myConstant)

	fmt.Println(num1, num2, num3, num4, num5)

}
