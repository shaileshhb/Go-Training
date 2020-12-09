package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	data := []byte("Hello From fileWrite.go")
	err := ioutil.WriteFile("file.txt", data, 0644)
	checkErr(err)
	fmt.Print("Successful")

	// append to current file contents
	file, err := os.OpenFile("file.txt", os.O_APPEND|os.O_WRONLY, 0644)
	checkErr(err)
	defer file.Close()

	if _, err := file.WriteString("\nappending to current file"); err != nil {
		log.Fatal(err)
	}

	// f, err := os.Create("file2.txt")
	// checkErr(err)
	// defer f.Close()

	// data2 := []byte("Hello From fileWrite.go")
	// bytesWritten, err := f.WriteString("Hello From fileWrite.go")
	// checkErr(err)
	// fmt.Println("Bytes Written: ", bytesWritten)
}

func checkErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
