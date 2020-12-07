package main

import "fmt"

func main() {
	ranks := map[string]int{"bronze": 3, "gold": 1, "silver": 2}

	for medal, rank := range ranks {
		fmt.Printf("The %s medal's rank is %d\n", medal, rank)
	}

}
