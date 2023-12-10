package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("Starting main...")
	// var fd, ld int
	var buf [117]byte // 116 characters per line, plus one for newline

	file, _ := os.Open("input.txt")
	defer file.Close()

	for y := 0; y < 220; y++ {
		file.Read(buf[:])
		fmt.Print(y, ": ")
		for _, b := range buf[10:39] {
			fmt.Printf("%c", b)
		}
		fmt.Print("\n")
	}
}
