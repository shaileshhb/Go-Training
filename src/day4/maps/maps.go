package main

import "fmt"

func main() {

	var nums map[string]int
	nums = make(map[string]int)
	nums["One"] = 1
	nums["Two"] = 2
	nums["Three"] = 3
	nums["Four"] = 4

	fmt.Println(nums)
	// fmt.Println(len(nums), cap(nums))
	fmt.Println(len(nums))

	// deleting value from map
	delete(nums, "Two")

	for key, value := range nums {
		fmt.Println(key, value)
	}

	// checking if key exists in the map
	// ok gives a bool value stating whether key is present or not
	var num, ok = nums["Three"]
	fmt.Println(num, ok)

	var num1, ok1 = nums["Two"]
	fmt.Println(num1, ok1)


}