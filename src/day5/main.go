package main

import (
	"fmt"
	"os"
	"strconv"
)

// don't forget newline before eof
const filename string = "test_input.txt"
const max int = 99

func isNumeric(b byte) bool {
	return b >= '0' && b <= '9'
}

func main() {
	// highestSeed := 0
	// lowestNum := max

	fileBuf := make([]byte, 1)
	digitBuf := make([]byte, 3)
	m := make(map[int]int)

	file, _ := os.Open(filename)
	defer file.Close()

	parseLine := func() []int {
		s := make([]int, 0)
		for {
			file.Read(fileBuf)
			if isNumeric(fileBuf[0]) {
				digitBuf = append(digitBuf, fileBuf[0])
			} else if fileBuf[0] == ' ' && len(digitBuf) > 0 {
				n, _ := strconv.Atoi(string(digitBuf))
				s = append(s, n)
				digitBuf = nil
			} else if fileBuf[0] == '\n' {
				n, _ := strconv.Atoi(string(digitBuf))
				s = append(s, n)
				digitBuf = nil
				break
			}
		}
		if len(s) > 1 {
			return s
		}
		return nil
	}

	seeds := parseLine()[1:]
	fmt.Println("Seeds: ", seeds)
	for i := 1; i < 33; i++ {
		v := parseLine()
		if v == nil {
			fmt.Println("eos")
			clear(m)
			continue
		}
		fmt.Printf("%v\n", v)
		for j := 0; j < v[2]; j++ {
			m[v[1]+j] = v[0] + j
		}
		fmt.Println(m)
	}
}
