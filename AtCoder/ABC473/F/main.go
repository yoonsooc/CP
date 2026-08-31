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

type Item struct {
	rangeSum     int
	minPrefixSum int
}

type SegTree struct {
	tree []Item
	n    int
}

func NewSegTree(arr []byte) *SegTree {
	n := len(arr)
	st := &SegTree{
		tree: make([]Item, 4*n),
		n:    n,
	}
	st.build(arr, 1, 0, n-1)
	return st
}

func merge(leftItem, rightItem Item, start, end int) Item {
	return Item{
		rangeSum:     leftItem.rangeSum + rightItem.rangeSum,
		minPrefixSum: min(leftItem.minPrefixSum, leftItem.rangeSum+rightItem.minPrefixSum),
	}
}

func (st *SegTree) build(arr []byte, node, start, end int) {
	if start == end {
		if arr[start] == 'A' {
			st.tree[node] = Item{
				rangeSum:     1,
				minPrefixSum: 1,
			}
		} else {
			st.tree[node] = Item{
				rangeSum:     -1,
				minPrefixSum: -1,
			}
		}
		return
	}

	mid := (start + end) / 2
	st.build(arr, 2*node, start, mid)
	st.build(arr, 2*node+1, mid+1, end)

	leftItem := st.tree[2*node]
	rightItem := st.tree[2*node+1]

	st.tree[node] = merge(leftItem, rightItem, start, end)
}

func (st *SegTree) query(start, end, node, left, right int) Item {
	outOfRange := right < start || end < left
	if outOfRange {
		return Item{
			rangeSum:     0,
			minPrefixSum: 0,
		}
	}

	included := left <= start && end <= right
	if included {
		return st.tree[node]
	}

	mid := (start + end) / 2
	leftSum := st.query(start, mid, 2*node, left, right)
	rightSum := st.query(mid+1, end, 2*node+1, left, right)
	return merge(leftSum, rightSum, start, end)
}

func (st *SegTree) update(start, end, node, idx int, val byte) {
	if idx < start || idx > end {
		return
	}

	if start == end {
		if val == 'A' {
			st.tree[node].rangeSum = 1
			st.tree[node].minPrefixSum = 1
		} else {
			st.tree[node].rangeSum = -1
			st.tree[node].minPrefixSum = -1
		}
		return
	}

	mid := (start + end) / 2
	if idx <= mid {
		st.update(start, mid, 2*node, idx, val)
	} else {
		st.update(mid+1, end, 2*node+1, idx, val)
	}
	st.tree[node] = merge(st.tree[2*node], st.tree[2*node+1], start, end)
}

func (st *SegTree) print(node, start, end int) {
	fmt.Println("NODE[", node, "]\t", start, "~", end, st.tree[node], string(S[start:end+1]))
	if start == end {
		return
	} else {
		mid := (start + end) / 2
		st.print(2*node, start, mid)
		st.print(2*node+1, mid+1, end)
	}
}

var N, Q int
var S []byte

func readInt() int {
	scanner.Scan()
	res, _ := strconv.Atoi(scanner.Text())
	return res
}

func change(sgt *SegTree) {
	i := readInt()
	scanner.Scan()
	c := scanner.Text()[0]
	S[i-1] = c
	sgt.update(0, N-1, 1, i-1, c)
	// sgt.print(1, 0, N-1)
}

func verify(sgt *SegTree) {
	l, r := readInt(), readInt()
	// sgt.print(1, 0, N-1)
	item := sgt.query(0, N-1, 1, l-1, r-1)
	// fmt.Println(item)
	if item.minPrefixSum >= 0 {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}

func main() {
	scanner = bufio.NewScanner(os.Stdin)
	writer = bufio.NewWriter(os.Stdout)

	defer writer.Flush()

	const maxCapacity = 5 * 1024 * 1024
	buf := make([]byte, maxCapacity)
	scanner.Buffer(buf, maxCapacity)
	scanner.Split(bufio.ScanWords)

	N = readInt()
	scanner.Scan()
	S = scanner.Bytes()
	Q = readInt()

	sgt := NewSegTree(S)
	// sgt.print(1, 0, N-1)
	for range Q {
		t := readInt()
		switch t {
		case 1:
			change(sgt)
		case 2:
			verify(sgt)
		}
	}
}
