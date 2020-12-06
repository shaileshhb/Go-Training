package main

import "fmt"

func main() {
	variadicFunc("string", 1, 2, 3, 4, 5)
	// arr := [...]int{1, 2, 3}
	// cant use array as the type of num in varidaic func is slice
	// [3]int and [...]int both cannot be passed to variadic func
	// variadicFunc("Array", arr...)

}

func variadicFunc(str string, num ...int) {
	fmt.Println(str)
	fmt.Printf("Type of num is %T\n", num)
	for i, num := range num {
		fmt.Println(i, num)
	}
}
