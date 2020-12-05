package main

import (
	"fmt"
)

func main() {

	for i := 0; i < 5; i++ {
		fmt.Println("hello", i)
	}

	j := 1

	for ; j < 5; j++ {
		fmt.Println(j)
	}
}
