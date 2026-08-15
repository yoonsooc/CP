package main

import (
	"bufio"
	"fmt"
	"math/big"
	"os"
	"strconv"
)

var (
	scanner *bufio.Scanner
	writer  *bufio.Writer
)

func readInt() int64 {
	scanner.Scan()
	res, _ := strconv.ParseInt(scanner.Text(), 10, 64)
	return res
}

// 원래 풀이(수직이등분선의 교차/일치 판정)를 big.Rat으로 구현한 버전.
// big.Rat은 값을 분자/분모 정수 쌍으로 정확히 들고 다니므로
// float64의 반올림 오차와 0/0(NaN) 문제가 모두 사라진다.
//
// 이등분선은 기울기(y = ax + b) 대신 일반형으로 세운다:
//   PQ의 수직이등분선 = { X : (X - M)·(Q-P) = 0 }   (M은 PQ의 중점)
//   => (qx-px)·x + (qy-py)·y = M·(Q-P)
// 이렇게 하면 수직/수평/일반 기울기의 세 갈래 분기가 필요 없다.
func solve() string {
	px, py := readInt(), readInt()
	qx, qy := readInt(), readInt()
	rx, ry := readInt(), readInt()
	sx, sy := readInt(), readInt()

	// PQ의 이등분선: a1*x + b1*y = c1
	a1 := big.NewRat(qx-px, 1)
	b1 := big.NewRat(qy-py, 1)
	c1 := new(big.Rat).Add(
		new(big.Rat).Mul(big.NewRat(px+qx, 2), a1),
		new(big.Rat).Mul(big.NewRat(py+qy, 2), b1))

	// RS의 이등분선: a2*x + b2*y = c2
	a2 := big.NewRat(sx-rx, 1)
	b2 := big.NewRat(sy-ry, 1)
	c2 := new(big.Rat).Add(
		new(big.Rat).Mul(big.NewRat(rx+sx, 2), a2),
		new(big.Rat).Mul(big.NewRat(ry+sy, 2), b2))

	// 평행이 아니면 교점(공통 중심)이 반드시 존재
	det := new(big.Rat).Sub(new(big.Rat).Mul(a1, b2), new(big.Rat).Mul(a2, b1))
	if det.Sign() != 0 {
		return "Yes"
	}

	// 평행이면 두 직선이 일치할 때만 Yes.
	// c1/c2 같은 나눗셈 비교 대신 곱셈 교차 비교를 쓴다.
	// (a1,b1,c1)과 (a2,b2,c2)가 비례하는지 확인 (0/0 상황이 원천적으로 없음)
	x := new(big.Rat).Sub(new(big.Rat).Mul(a1, c2), new(big.Rat).Mul(a2, c1))
	y := new(big.Rat).Sub(new(big.Rat).Mul(b1, c2), new(big.Rat).Mul(b2, c1))
	if x.Sign() == 0 && y.Sign() == 0 {
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
