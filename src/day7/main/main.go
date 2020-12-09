package main

import (
	"day7/customer"
	"day7/order"
	"fmt"
)

func main() {

	ch := make(chan uint)

	var custOneOrder []order.Orders
	custOne := customer.SetCustomerInfo(101, "Sam", "Thomas", &custOneOrder)

	go custOne.PlaceOrder(2, 4000, ch)
	go custOne.PlaceOrder(1, 1000, ch)

	fmt.Println("Order Successfully Placed with order id", <-ch)
	fmt.Println("Order Successfully Placed with order id", <-ch)

}
