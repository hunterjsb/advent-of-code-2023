package main

import (
	"fmt"
	"io"
	"os"
)

// warning! make sure there is an empty line at the bottom of input.txt!

const filename string = "input.txt"

var TextDigits = [10]string{"zero", "one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}

var b [1]byte

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func isNumeric(b byte) bool {
	return b >= '0' && b <= '9' // '0' to '9' in ASCII
}

func convertDigits(a, b byte) int {
	firstDigit := a - '0'
	lastDigit := b - '0'
	return int(firstDigit)*10 + int(lastDigit)
}

func main() {
	// open file
	file, err := os.Open(filename)
	check(err)
	defer file.Close()

	// iterate over the bytes
	var fn, ln byte
	sum := 0
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
		} else if b[0] == '\n' {
			n := convertDigits(fn, ln)
			sum += n
			fn, ln = 0, 0
		}
	}
	fmt.Println("sum", sum)
}
