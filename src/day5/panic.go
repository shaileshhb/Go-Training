package main

import "fmt"

func main() {
	fmt.Println("Main Start")

	// recover from panic, used mainly for closing resources
	defer func() {
		details := recover()
		fmt.Println("Details:", details)
	}()
	//panic
	panic("Paniced")

	fmt.Print("Main End")
}
