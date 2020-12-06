package main

import "fmt"

func main() {
	myMap := map[int]string{
		1: "a",
		2: "b",
		3: "c",
		4: "d",
	}

	newMap := checkMap(myMap)
	fmt.Println(myMap)
	fmt.Println(newMap)
}

func checkMap(myMap map[int]string) map[int]string {
	myMap[2] = "e"
	myMap[4] = "f"
	return myMap
}