package main

import (
	"bufio"
	"fmt"
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

func solve() string {
	px := readInt()
	py := readInt()
	qx := readInt()
	qy := readInt()
	rx := readInt()
	ry := readInt()
	sx := readInt()
	sy := readInt()

	cross := (qx-px)*(sy-ry) - (sx-rx)*(qy-py)
	if cross != 0 {
		return "Yes"
	}
	dot := (rx+sx-px-qx)*(qx-px) + (ry+sy-py-qy)*(qy-py)
	if dot == 0 {
		return "Yes"
	}
	return "No"
}

func main() {
	scanner = bufio.NewScanner(os.Stdin)
	writer = bufio.NewWriter(os.Stdout)

	defer writer.Flush()

	const maxCapacity = 5 * 1024 * 1024
	buf := make([]byte, maxCapacity)
	scanner.Buffer(buf, maxCapacity)
	scanner.Split(bufio.ScanWords)

	T := readInt()
	for range T {
		fmt.Fprintln(writer, solve())
	}
}
