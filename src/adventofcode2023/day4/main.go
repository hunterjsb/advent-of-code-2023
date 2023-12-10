package main

import (
	"fmt"
	"os"
	"slices"
)

func toInt(a byte, b byte) int {
	if a < '0' {
		return int(b - '0')
	}
	return int(a-'0')*10 + int(b-'0')
}

func main() {
	fmt.Println("Starting main...")
	var buf [117]byte // 116 characters per line, plus one for newline
	var sum, matches int
	nums := make([]int, 0)

	file, _ := os.Open("input.txt")
	defer file.Close()

	for y := 0; y < 220; y++ {
		file.Read(buf[:])
		fmt.Print(y, ": ")
		for i := 10; i < 39; i += 3 {
			nums = append(nums, toInt(buf[i], buf[i+1]))
		}
		fmt.Print(nums, " ... ")
		for i := 42; i < 115; i += 3 {
			fmt.Printf("%c%c ", buf[i], buf[i+1])
			n := toInt(buf[i], buf[i+1])
			if slices.Contains(nums, n) {
				matches++
			}
		}
		fmt.Print("\n")
		if matches > 0 {
			sum += 1 << (matches - 1)
		}
		nums = nil
		matches = 0
	}
	fmt.Println("SUM: ", sum)
}
