package main

import (
	"fmt"
	"io"
	"os"
)

const filename string = "input.txt"

var b [1]byte

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func isNumeric(b byte) bool {
	return b >= '0' && b <= '9' // '0' to '9' in ASCII
}

func main() {
	// open file
	file, err := os.Open(filename)
	check(err)
	defer file.Close()

	// iterate over the bytes
	var fn, ln byte
	for {
		_, err := file.Read(b[:])
		if err == io.EOF {
			break
		}

		if isNumeric(b[0]) {
			if fn == 0 {
				fn = b[0]
			}
			ln = b[0]
		}
	}
	fmt.Println(fn, ln)
}
