package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(concatenateName("sam", "thomas"))
}

func concatenateName(firstName, lastName string) string {
	return strings.ToUpper(firstName + " " + lastName)
}
