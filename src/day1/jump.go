package main

import (
	"fmt"
)

func main() {
	var userNum int
	// label for go to
test:
	fmt.Print("Number: ")
	fmt.Scanln(&userNum)
	fmt.Println(userNum)

	if userNum > 100 {
		fmt.Println("enter another num")
		goto test
	} else {
		fmt.Println("Done")
	}

	for i := 0; i < 10; i++ {
		if i == 4 {
			continue
		}
		if i == 7 {
			break
		}

		fmt.Println("hello", i)
	}

}
