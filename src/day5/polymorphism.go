package main

import "fmt"

func main() {

	bmw := &car{brandName: "BMW"}
	pulsar := &bike{brandName: "Pulsar"}
	busName := &bus{brandName: "Bus Name"}

	isMoving(bmw)
	isMoving(pulsar)
	isMoving(busName)

}

func isMoving(m movable) {
	fmt.Println(m.start())
}

func (bi *bike) start() string {
	return bi.brandName + " is moving"
}

func (bu *bus) start() string {
	return bu.brandName + " is moving"
}

func (c *car) start() string {
	return c.brandName + " is moving"
}

type movable interface {
	start() string
}

type car struct {
	brandName string
}

type bus struct {
	brandName string
}

type bike struct {
	brandName string
}
