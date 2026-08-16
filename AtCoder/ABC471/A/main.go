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

func main() {
	scanner = bufio.NewScanner(os.Stdin)
	writer = bufio.NewWriter(os.Stdout)

	defer writer.Flush()

	const maxCapacity = 5 * 1024 * 1024
	buf := make([]byte, maxCapacity)
	scanner.Buffer(buf, maxCapacity)
	scanner.Split(bufio.ScanWords)

	A := readInt()
	B := readInt()

	if A+B == 9 || A-B == 9 || A*B == 9 || float32(A)/float32(B) == 9.0 {
		fmt.Fprintln(writer, "Nine")
	} else {
		fmt.Fprintln(writer, "Nein")
	}
}
