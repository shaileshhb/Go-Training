package hello

import "fmt"

func main() {
	fmt.Println("Hello world")
	testForGo()
	TestForGo()
}

// private function uses camelCasing
func testForGo() {
	fmt.Println("testForGo")
}

func TestForGo() {
	// public function in go
	fmt.Println("TestForGo")
}
