package main

import "fmt"

func main() {

	num := 50

	defer f2(num)

	num = 100

	defer f1()
	fmt.Println("Inside main. Number is:", num)
}

func f1() {
	defer fmt.Println("Defer in f1()")
	fmt.Println("Inside f1()")
}

func f2(num int) {
	fmt.Println("Inside f2(). Number is:", num)
}
