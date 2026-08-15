package main

import (
	"bufio"
	"fmt"
	"math"
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

func check(r float64, N, K int, oCnt []int) bool {
	minPrefix := math.Inf(1)
	p := 0
	for j := 1; j <= N; j++ {
		c := oCnt[j] - K
		if c < 0 {
			continue
		}
		for p <= N && oCnt[p] <= c {
			v := float64(oCnt[p]) - r*float64(p)
			minPrefix = min(minPrefix, v)
			p++
		}
		if float64(oCnt[j])-r*float64(j) >= minPrefix {
			return true
		}
	}
	return false
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
	K := readInt()

	scanner.Scan()
	S := scanner.Text()

	oCnt := make([]int, N+1)
	for i := range N {
		oCnt[i+1] = oCnt[i]
		if S[i] == 'o' {
			oCnt[i+1]++
		}
	}

	lo, hi := 0.0, 1.0
	for range 60 {
		mid := (lo + hi) / 2
		if check(mid, N, K, oCnt) {
			lo = mid
		} else {
			hi = mid
		}
	}

	fmt.Fprintf(writer, "%.10f\n", lo)
}
