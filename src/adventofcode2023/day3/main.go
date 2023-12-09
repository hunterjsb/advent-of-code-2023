package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	buf := make([]byte, 140) // Create a buffer of 140 bytes
	scanner.Buffer(buf, 140) // Set the buffer size to 140 bytes

	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println(line) // Process the line
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading from file:", err)
	}
}
