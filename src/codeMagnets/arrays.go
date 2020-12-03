package main

import (
	"fmt"
)

func main() {
	var arr [4]int = [4]int{1, 2, 3, 4}

	// for i := 0; i < len(arr); i++ {
	// 	fmt.Println(arr[i])
	// }

	ar := [3]int{1, 2, 3}

	// for i := 0; i < len(ar); i++ {
	// 	fmt.Println(arr[i])
	// }

	fmt.Println(ar)

	// fo...range loop for arrays
	for index, value := range arr {
		fmt.Println(index, value, " ")
	}

}
