package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
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

func findMax(m map[string]int) int {
	maxVal := 0
	for _, v := range m {
		if v > maxVal {
			maxVal = v
		}
	}
	return maxVal
}

func main() {
	scanner = bufio.NewScanner(os.Stdin)
	writer = bufio.NewWriter(os.Stdout)

	defer writer.Flush()

	const maxCapacity = 5 * 1024 * 1024
	buf := make([]byte, maxCapacity)
	scanner.Buffer(buf, maxCapacity)
	scanner.Split(bufio.ScanWords)

	N := readInt()
	m := map[string]int{}
	for range N {
		scanner.Scan()
		S := strings.ToUpper(scanner.Text())
		m[S]++
	}
	fmt.Fprintln(writer, findMax(m))
}
