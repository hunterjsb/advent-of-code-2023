package main

import (
	"fmt"
	"os"
	"strconv"
)

// don't forget newline before eof
const filename string = "test_input.txt"

func isNumeric(b byte) bool {
	return b >= '0' && b <= '9'
}

func main() {
	buf := make([]byte, 1)
	num := make([]byte, 0)
	// seeds := make([]int, 0)

	file, _ := os.Open(filename)
	defer file.Close()

	for i := 0; i < 33; i++ {
		for {
			file.Read(buf)
			if isNumeric(buf[0]) {
				num = append(num, buf[0])
			} else if buf[0] == 32 && len(num) > 0 {
				n, _ := strconv.Atoi(string(num))
				fmt.Printf(" [[%v]] ", n)
				num = nil
			} else if buf[0] == 10 {
				n, _ := strconv.Atoi(string(num))
				fmt.Printf(" [{%v}]\n", n)
				num = nil
				break
			}
		}
	}
}
