package main

import (
	"fmt"
	"keyboard"
	"log"
)

func main() {

	fmt.Print("Enter a number: ")
	grade, err := keyboard.GetFloat()
	if err != nil {
		log.Fatal(err)
	}

	var status string
	if grade >= 60 {
		status = "passed"
	} else {
		status = "failed"
	}

	fmt.Println(status)
}