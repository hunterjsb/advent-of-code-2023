package main

import (
	"fmt"
	"os"
	"slices"
)

// toInt converts a pair of ASCII characters representing digits
// into their integer value. If the first character is not a digit,
// it returns the value of the second character as an integer.
func toInt(a byte, b byte) int {
	if a < '0' {
		return int(b - '0')
	}
	return int(a-'0')*10 + int(b-'0')
}

func main() {
	var buf [117]byte // 116 characters per line, plus one for newline
	var sum, matches int
	hand := make([]int, 0) // Slice to store integer values for each hand

	file, _ := os.Open("input.txt")
	defer file.Close()

	for y := 0; y < 220; y++ {
		file.Read(buf[:])

		// Parse the first part of the line to get hand values
		for i := 10; i < 39; i += 3 {
			hand = append(hand, toInt(buf[i], buf[i+1]))
		}

		// Check for matches in the second part of the line
		for i := 42; i < 115; i += 3 {
			n := toInt(buf[i], buf[i+1])
			if slices.Contains(hand, n) {
				matches++
			}
		}

		// If there are matches, add to sum based on the number of matches
		if matches > 0 {
			sum += 1 << (matches - 1) // Double the sum for each match
		}

		// Reset hand and matches for the next line
		hand = nil
		matches = 0
	}
	fmt.Println("SUM: ", sum) // Print the final sum
}
