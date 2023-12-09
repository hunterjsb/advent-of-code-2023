package main

import (
	"fmt"
	"io"
	"os"
)

// warning! make sure there is an empty line at the bottom of input.txt!

const filename string = "input.txt"

var TextDigits = [10]string{"zero", "one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func isDigit(b byte) bool {
	return b >= '0' && b <= '9'
}

func convertDigits(a, b byte) int {
	firstDigit := a - '0' // ascii to int, same as a - 48
	lastDigit := b - '0'
	return int(firstDigit)*10 + int(lastDigit)
}

func setN(first, n byte) (byte, byte) {
	if first == 0 {
		first = n
	}
	return first, n
}

func main() {
	var curNode *Node
	//var prefix []rune
	var buf [1]byte

	// open file
	file, err := os.Open(filename)
	check(err)
	defer file.Close()

	// iterate over the bytes
	var fn, ln byte // first num, last num per line
	sum := 0
	for {
		// read next byte and check for eof
		_, err := file.Read(buf[:])
		b := buf[0]
		fmt.Printf("%c! %p\n", b, curNode)
		if err == io.EOF {
			break
		}

		// check if it's a digit
		if isDigit(b) {
			fn, ln = setN(fn, ln)
			fmt.Printf("is digit %c %v %v#\n", b, fn, ln)
		}

		if b == '\n' {
			sum += convertDigits(fn, ln)
			fmt.Printf("summing %v to %v\n", convertDigits(fn, ln), sum)
		}
	}

	fmt.Println("sum", sum)
}
