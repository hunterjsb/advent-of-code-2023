package main

import (
	"fmt"
	"os"
)

const ChunkSize int = 512

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
	chunks := size / ChunkSize
	remainder := size % ChunkSize
	fmt.Printf("%v bytes | %v chunks | Remainder %v\n", size, chunks, remainder)

	// iterate over the bytes
	buf := make([]byte, ChunkSize)
	for i := 0; i < chunks; i++ {
		n, err := f.Read(buf)
		check(err)
		fmt.Println("\n", n, "buff ", buf)
	}
}
