package main

import "fmt"

func main() {

	var circle = &circle{
		radius: 4.5,
	}
	passInterface(circle)
}

type shape interface {
	area() float64
	perimeter() float64
}

type circle struct {
	radius float64
}

func (c *circle)area() float64 {
	return 22/7 * (c.radius * c.radius)
}

func (c *circle)perimeter() float64 {
	return 22/7 * c.radius
}


func passInterface(sh shape) {
	fmt.Println("Area: ",sh.area())
	fmt.Println("Perimeter: ",sh.perimeter())
	
}