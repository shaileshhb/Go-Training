package order

func SetOrder(id, quantity, cost int) *Orders {
	return &Orders{
		id:       id,
		quantity: quantity,
		cost:     cost,
	}
}

func (o *Orders) ID() int {
	return o.id
}

func (o *Orders) Quantity() int {
	return o.quantity
}

func (o *Orders) Cost() int {
	return o.cost
}

// Orders will have orders info
type Orders struct {
	id       int
	quantity int
	cost     int
}
