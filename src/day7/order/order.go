package order

func SetOrder(id uint, quantity, cost int) *Orders {
	return &Orders{
		id:       id,
		quantity: quantity,
		cost:     cost,
	}
}

func (o *Orders) ID() uint {
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
	id       uint
	quantity int
	cost     int
}
