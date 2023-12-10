package main

import (
	"fmt"
	"os"
	"strconv"
)

type point struct {
	x, y int
}

func (p point) adjacent() []point {
	adj := make([]point, 0, 8) // Initialize with a capacity of 8 for all adjacent points

	// Loop through the possible offsets for x and y
	for dx := -1; dx <= 1; dx++ {
		for dy := -1; dy <= 1; dy++ {
			if dx == 0 && dy == 0 {
				continue // Skip the current point itself
			}

			newX := p.x + dx
			newY := p.y + dy

			// Check if the new point is within bounds
			if newX >= 0 && newX < 140 && newY >= 0 && newY < 140 {
				adj = append(adj, point{newX, newY})
			}
		}
	}

	return adj
}

func isDigit(b byte) bool {
	if b >= '0' && b <= '9' {
		return true
	}
	return false
}

func isSymbol(b byte) bool {
	if b != '.' && !isDigit(b) {
		return true
	}
	return false
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	const filename string = "input.txt"
	var buf [141]byte // 140 characters per line, 141 with newline

	symbols := make(map[point]byte)
	lastNumbers := make(map[*int][]point, 0)
	digits := make([]byte, 0)
	points := make([]point, 0)

	sum := 0

	// open file
	file, err := os.Open(filename)
	check(err)
	defer file.Close()

	// iterate line-by-line
	for y := 0; y < 140; y++ {
		file.Read(buf[:])
		for x, b := range buf[:len(buf)-1] { // minus one cuz last char is newline
			if isSymbol(b) { // keep track of the location of symbols so we can see if any are adjacent to numbers
				symbols[point{x, y}] = b
			}

			if isDigit(b) { // can't be `else if` because symbol can delimit numbers
				digits = append(digits, b)
				points = append(points, point{x, y})
			} else if len(digits) > 0 { // current byte is not digit but we have digits in the slice
				n, _ := strconv.Atoi(string(digits))
				pointsCopy := make([]point, len(points))
				copy(pointsCopy, points)
				lastNumbers[&n] = pointsCopy
				digits = digits[:0]
				points = points[:0]
			}
		}
	}
	for num, pts := range lastNumbers {
	numLoop:
		for _, pt := range pts {
			for _, adjPt := range pt.adjacent() {
				//fmt.Printf("%v is adj to %v\n", adjPt, pt)
				if _, exists := symbols[adjPt]; exists {
					sum += *num
					fmt.Printf("adding %v to %v \n", num, sum)
					break numLoop
				}
			}
		}
	}
	fmt.Printf("SUM = %v\n", sum)
}
