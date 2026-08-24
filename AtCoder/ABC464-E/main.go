package main

import (
	"bufio"
	"os"
	"strconv"
)

var (
	scanner *bufio.Scanner
	writer  *bufio.Writer
)

func readInt() int {
	scanner.Scan()
	res, _ := strconv.Atoi(scanner.Text())
	return res
}

func main() {
	scanner = bufio.NewScanner(os.Stdin)
	writer = bufio.NewWriter(os.Stdout)

	defer writer.Flush()

	const maxCapacity = 5 * 1024 * 1024
	buf := make([]byte, maxCapacity)
	scanner.Buffer(buf, maxCapacity)
	scanner.Split(bufio.ScanWords)

	H := readInt()
	W := readInt()

	grid := make([][]int, H)
	for i := range H {
		grid[i] = make([]int, W)
	}
	Q := readInt()
	Xs := make([]byte, Q+1)
	Xs[0] = 'A'
	for xi := range Q {
		i := readInt() - 1
		j := readInt() - 1
		scanner.Scan()
		X := byte(scanner.Text()[0])
		Xs[xi+1] = X
		grid[i][j] = xi + 1
	}
	for row := H - 1; row >= 0; row-- {
		for col := W - 1; col >= 0; col-- {
			if row != 0 {
				grid[row-1][col] = max(grid[row][col], grid[row-1][col])
			}
			if col != 0 {
				grid[row][col-1] = max(grid[row][col], grid[row][col-1])
			}
		}
	}

	for r := range H {
		row := make([]byte, len(grid[r]))
		for i, v := range grid[r] {
			row[i] = Xs[v]
		}
		writer.WriteString(string(row))
		writer.WriteString("\n")
	}

}
