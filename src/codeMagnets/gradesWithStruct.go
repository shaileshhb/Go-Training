package main

import "fmt"

func main() {

	var s student
	s.name = "Alonzo Cole"
	s.grade = 92.3
	printInfo(s)

}

func printInfo(s student) {
	fmt.Println("Name: ", s.name)
	fmt.Printf("Grade: %0.1f\n", s.grade)
}

type student struct {
	name  string
	grade float64
}
