package main

import (
	"bufio"
	"container/heap"
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

type PriorityQueue []int

func (pq PriorityQueue) Len() int {
	return len(pq)
}
func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i] > pq[j]
}
func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}
func (pq *PriorityQueue) Push(i any) {
	*pq = append(*pq, i.(int))
}
func (pq *PriorityQueue) Pop() any {
	old := *pq
	n := len(old)
	item := old[n-1]
	*pq = old[:n-1]

	return item
}

func main() {
	scanner = bufio.NewScanner(os.Stdin)
	writer = bufio.NewWriter(os.Stdout)

	defer writer.Flush()

	const maxCapacity = 5 * 1024 * 1024
	buf := make([]byte, maxCapacity)
	scanner.Buffer(buf, maxCapacity)
	scanner.Split(bufio.ScanWords)

	Q := readInt()
	V := readInt()

	pq := PriorityQueue{}
	heap.Init(&pq)

	for range Q {
		q := readInt()
		t := readInt()
		if q == 1 {
			w := readInt()
			heap.Push(&pq, w-t)
		} else {
			if pq.Len() == 0 {
				fmt.Fprintln(writer, -1)
				continue
			}
			u := heap.Pop(&pq).(int)
			fmt.Fprintln(writer, min(V, u+t))
		}
	}
}
