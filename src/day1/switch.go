package main

import (
	"fmt"
)

func main() {
	cgpa := 8

	switch cgpa {
	case 10:
		fmt.Println("10 CGPA")
	case 9, 8:
		fmt.Println("9 cgpa")
	default:
		fmt.Println("cgpa not found")
	}

	//break stmt not required
}
