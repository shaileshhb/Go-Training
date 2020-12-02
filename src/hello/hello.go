package hello

import "fmt"

func main() {
	fmt.Println("Hello world")
}

// private function uses camelCasing
func testForGo() {
	fmt.Println("testForGo")
}

// public function in go
func TestForGo() {
	fmt.Println("TestForGo")
}
