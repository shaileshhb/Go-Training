package main

import (
	"fmt"
	"keyboard/datafile"
	"log"
)

func main() {

	lines, err := datafile.GetStrings("voters.txt")

	if err != nil {
		log.Fatal(err)
	}

	var names []string
	var count []int
	for _, line := range lines {
		matched := false
		for i, name := range names {
			if name == line {
				count[i]++
				matched = true
			}
		}
		if matched == false {
			names = append(names, line)
			count = append(count, 1)
		}
	}

	for i, name := range names {
		fmt.Printf("%s: %d\n", name, count[i])
	}

	//map[keyType]valueType
	// var votes map[string]int
	// votes = make(map[string]int)

		
}
