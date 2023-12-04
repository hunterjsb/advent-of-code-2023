package main

import (
	"fmt"
	"io"
	"os"
)

var buf [1]byte

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func getFile(filename string) *os.File {
	file, err := os.Open(filename)
	check(err)
	return file
}

func main() {
	// open file
	f := getFile("input.txt")
	defer f.Close()
	info, err := f.Stat()
	check(err)

	// get the size
	size := int(info.Size())
	fmt.Printf("%v bytes", size)

	// iterate over the bytes
	for i := 0; i < size; i++ {
		_, err := f.Read(buf[:])
		if err == io.EOF {
			break // End of file reached
		}
		if err != nil {
			check(err) // Handle other errors
		}
		fmt.Println("buff ", buf)
	}
}
