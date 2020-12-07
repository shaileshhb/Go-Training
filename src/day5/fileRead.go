package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func main() {

	content, err := ioutil.ReadFile("file.txt")
	checkError(err)
	fmt.Println("Content: ", string(content))

	f, err := os.Open("file.txt")
	defer f.Close()
	checkError(err)

	bucket := make([]byte, 10)
	bytesRead, err := f.Read(bucket)
	checkError(err)

	fmt.Println("Content:", string(bucket), "\nbytes read:", bytesRead)
}

func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
