package main

import (
	"fmt"
)

func main() {

	quotient, err := divide(4, 0)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(quotient)
}

type myError struct {
	s string
}

func (err *myError) Error() string {
	return err.s
}

func divide(numerator, denominator int) (int, error) {

	if denominator == 0 {
		return 0, &myError{s: "Cannot Divide by Zero"}
	}
	return numerator / denominator, nil
}
