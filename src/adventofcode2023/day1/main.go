package main

import (
	"fmt"
	"os"
)

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
	f := getFile("day1/input.txt")
	defer f.Close()
	info, err := f.Stat()
	check(err)

	// get the size
	size := int(info.Size())
	fmt.Printf("Size: %v KB\n", float64(size)/1000.0)

	for i := 0; i < size; i++ {
		fmt.Println(i)
	}
}
