package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now().Hour()
	if now <= 12 {
		fmt.Println("Morning")
	} else if now <= 18 {
		fmt.Println("Afternoon")
	} else {
		fmt.Println("Night")
	}
}
