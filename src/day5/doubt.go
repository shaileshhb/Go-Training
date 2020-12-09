package main

import "fmt"

func main() {

	funcA(&circle{radius: 12.2})
	funcB(&circle{radius: 12.2})

}

func (c1 *circle) area() float64 {
	return 22 / 7 * (c1.radius * c1.radius)
}

type shape1 interface {
	area() float64
}

type shape2 interface {
	area() float64
}

type circle struct {
	radius float64
}

func funcA(sh shape1) {
	fmt.Println(sh.area())
}

func funcB(sh shape2) {
	fmt.Println(sh.area())
}
