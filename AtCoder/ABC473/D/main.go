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

func joinAndPrintIntSlices(nums []int) {
	var b strings.Builder
	for i, s := range nums {
		if i > 0 {
			b.WriteString(" ")
		}
		fmt.Fprint(&b, s)
	}
	fmt.Fprintln(writer, b.String())
}

func dfs(i, sum int, As []int) {
	isLastCoeff := i == N-1
	hasNoRemainder := (K-sum)%N == 0
	if isLastCoeff {
		if hasNoRemainder {
			As[i] = (K - sum) / N
			joinAndPrintIntSlices(As)
		}
		return
	}
	for x := range (K-sum)/(i+1) + 1 {
		As[i] = x
		dfs(i+1, sum+x*(i+1), As)
	}

}

var N, K int

func main() {
	scanner = bufio.NewScanner(os.Stdin)
	writer = bufio.NewWriter(os.Stdout)

	defer writer.Flush()

	const maxCapacity = 5 * 1024 * 1024
	buf := make([]byte, maxCapacity)
	scanner.Buffer(buf, maxCapacity)
	scanner.Split(bufio.ScanWords)

	N, K = readInt(), readInt()
	As := make([]int, N)
	// 1 * A1 + 2 * A2 + 3* A3 + ... + N * AN = K
	dfs(0, 0, As)
}
