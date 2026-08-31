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

var N, K int

type Item struct {
	maxPairs int
	index    int
}

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
	partSums := make([]int, N)
	maxPairs := make([]Item, N)
	lastRemainderPos := make(map[int]int, N)
	lastRemainderPos[0] = -1

	for i := range N {
		v := readInt()
		As[i] = v
		partSums[i] = v
		if i > 0 {
			partSums[i] += partSums[i-1]
		}
	}

	for i := range N {
		rem := partSums[i] % K
		if i == 0 {
			maxPairs[i].index = -1
			if rem == 0 {
				maxPairs[i] = Item{1, 0}
			}
			lastRemainderPos[rem] = i
			continue
		}

		mp := maxPairs[i-1]
		startIndex := mp.index + 1

		if lrp, exist := lastRemainderPos[rem]; exist && lrp >= startIndex-1 {
			maxPairs[i] = Item{mp.maxPairs + 1, i}
		} else {
			maxPairs[i] = Item{mp.maxPairs, mp.index}
		}
		lastRemainderPos[rem] = i
	}

	fmt.Fprintln(writer, maxPairs[N-1].maxPairs)
}
