package main

import (
	"fmt"
	"log"
)

func main() {

	// take ip using bufio
	// fmt.Print("Enter String: ")
	// reader := bufio.NewReader(os.Stdin)
	// input, err := reader.ReadString('\n')
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// input = strings.TrimSpace(input)
	// fmt.Print("Input:", input)

	// take ip using fmt.Scan
	fmt.Print("Enter String: ")
	var number int
	_, err := fmt.Scan(&number)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Print("Number: ", number)
}
