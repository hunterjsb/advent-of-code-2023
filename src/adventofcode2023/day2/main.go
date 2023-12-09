package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	maxima := map[rune]int{
		'r': 12,
		'g': 13,
		'b': 14,
	}

	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	sum := 0
	hasInvalid := false
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.Split(scanner.Text(), " ")
		id, _ := strconv.Atoi(strings.TrimRight(line[1], ":"))
		for i := 2; i < len(line)-1; i += 2 {
			strN, color := line[i], line[i+1]
			n, _ := strconv.Atoi(strN)
			max := maxima[rune(color[0])]
			if n > max {
				hasInvalid = true
			}
		}
		if !hasInvalid {
			sum += id
		}
		hasInvalid = false
	}

	fmt.Printf("SUM: %v\n", sum)

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading input:", err)
	}
}
