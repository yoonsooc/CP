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

const mod = 998244353

func modpow(base, exp int) int {
	res := 1
	base %= mod
	for exp > 0 {
		if exp&1 == 1 {
			res = res * base % mod
		}
		base = base * base % mod
		exp >>= 1
	}
	return res
}

func comb(n, k int) int {
	a := 1
	for i := n; i > k; i-- {
		a *= i
		a %= mod
	}
	b := 1
	for i := (n - k); i > 1; i-- {
		b *= i
		b %= mod
	}
	return a * modpow(b, mod-2) % mod
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

	// C(n-1, k-1) * sum of Ai^2 + C(n-2, k-2) * sum of Ai*Aj(i!=j)
	squareCoeff := comb(N-1, K-1)
	timesCoeff := comb(N-2, K-2)
	sum := 0
	sumOfSquares := 0
	for range N {
		A := readInt()
		sum += (A % mod)
		sumOfSquares += (A % mod) * (A % mod) % mod
		sumOfSquares %= mod
	}
	squareOfSums := ((sum % mod) * (sum % mod)) % mod
	timesSum := (squareOfSums - sumOfSquares + mod) % mod
	fmt.Fprintln(writer, ((squareCoeff*sumOfSquares)%mod+(timesCoeff*timesSum)%mod)%mod)

}
