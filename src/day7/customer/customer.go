package customer

import (
	"day7/order"
	"math/rand"
)

func SetCustomerInfo(id int, firstName, lastName string, orders *[]order.Orders) *Customer {

	return &Customer{
		id:        id,
		firstName: firstName,
		lastName:  lastName,
		orders:    orders,
	}

}

func (c *Customer) PlaceOrder(quantity, cost int, ch chan uint) {

	var orderID uint = uint(rand.Intn(50))

	order := order.SetOrder(orderID, quantity, cost)
	*c.orders = append(*c.orders, *order)

	ch <- order.ID()
}

func (c *Customer) ID() int {
	return c.id
}

func (c *Customer) Name() string {
	return c.firstName + " " + c.lastName
}

func (c *Customer) Orders() *[]order.Orders {
	return c.orders
}

// Customer will have customer info
type Customer struct {
	id        int
	firstName string
	lastName  string
	orders    *[]order.Orders
}
