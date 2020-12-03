package main

import "fmt"

func main() {

	var price int = 100
	fmt.Println("Price is ", price)

	var taxRate float64 = 0.08
	var tax float64 = taxRate * float64(price)
	fmt.Println("Tax is", tax)

	var total float64 = float64(price) + tax

	var availableFund int = 120
	fmt.Println(availableFund, "available")
	fmt.Println("Within budget", availableFund >= total)
}
