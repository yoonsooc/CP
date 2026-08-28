package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
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

type Bisect struct {
	arr []int
}

func NewBisect() *Bisect {
	return &Bisect{
		make([]int, 0),
	}
}

func bisectLeft(a []int, v int) int {
	l, r := 0, len(a)
	for l < r {
		m := (l + r) / 2
		if a[m] < v {
			l = m + 1
		} else {
			r = m
		}
	}
	return l
}

func (b *Bisect) add(v int) {
	pos := bisectLeft(b.arr, v)
	if pos == len(b.arr) {
		b.arr = append(b.arr, v)
	} else {
		b.arr[pos] = v
	}
}

type Line struct {
	A int
	B int
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
	lines := make([]Line, N)
	for i := range N {
		A, B := readInt(), readInt()
		lines[i] = Line{A, B}
	}
	sort.Slice(lines, func(i, j int) bool {
		if lines[i].A == lines[j].A {
			return lines[i].B > lines[j].B
		}
		return lines[i].A < lines[j].A
	})
	bs := NewBisect()
	for i := range N {
		curB := lines[i].B
		bs.add(curB)
	}

	fmt.Fprintln(writer, len(bs.arr))

}
