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

	f, err := os.Create("file2.txt")
	checkErr(err)
	defer f.Close()

	data2 := []byte("Hello From fileWrite.go")
	bytesWritten, err := f.Write(data2)
	checkErr(err)
	fmt.Println("Bytes Written: ", bytesWritten)
}

func checkErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
