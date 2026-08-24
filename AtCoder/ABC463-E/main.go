package main

import (
	"bufio"
	"container/heap"
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

type Edge struct {
	weight  int
	current int
}

type PQ []Edge

func (pq PQ) Len() int {
	return len(pq)
}
func (pq PQ) Less(i, j int) bool {
	return pq[i].weight < pq[j].weight
}
func (pq PQ) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}
func (pq *PQ) Push(x any) {
	*pq = append(*pq, x.(Edge))
}
func (pq *PQ) Pop() any {
	old := *pq
	n := len(old)
	top := old[n-1]
	*pq = old[:n-1]
	return top
}

func main() {
	scanner = bufio.NewScanner(os.Stdin)
	writer = bufio.NewWriter(os.Stdout)

	defer writer.Flush()

	const K = 1024
	const maxCapacity = 5 * K * K
	buf := make([]byte, maxCapacity)
	scanner.Buffer(buf, maxCapacity)
	scanner.Split(bufio.ScanWords)

	N := readInt()
	M := readInt()
	Y := readInt()

	bidirGraph := make([]map[int]int, N+1+1+1)
	for range M {
		u := readInt()
		v := readInt()
		T := readInt()

		if bidirGraph[u] == nil {
			bidirGraph[u] = make(map[int]int)
		}
		if bidirGraph[v] == nil {
			bidirGraph[v] = make(map[int]int)
		}

		if bidirGraph[u][v] == 0 || bidirGraph[u][v] > T {
			bidirGraph[u][v] = T
		}
		if bidirGraph[v][u] == 0 || bidirGraph[v][u] > T {
			bidirGraph[v][u] = T
		}

	}

	bidirGraph[N+1] = make(map[int]int)
	bidirGraph[N+2] = make(map[int]int)
	for i := range N {
		X := readInt()
		if bidirGraph[i+1] == nil {
			bidirGraph[i+1] = make(map[int]int)
		}
		bidirGraph[i+1][N+1] = X
		bidirGraph[N+2][i+1] = X
	}
	bidirGraph[N+1][N+2] = Y
	// fmt.Println(bidirGraph)

	// dijkstra
	shortestDistancesFromStart := make([]int, N+1+1+1)
	for i := range N + 1 + 1 + 1 {
		shortestDistancesFromStart[i] = math.MaxInt64
	}
	shortestDistancesFromStart[1] = 0

	edgeQueue := PQ{}
	heap.Init(&edgeQueue)
	heap.Push(&edgeQueue, Edge{0, 1})

	for edgeQueue.Len() > 0 {
		edge := heap.Pop(&edgeQueue).(Edge)
		currentWeight := edge.weight
		current := edge.current

		if shortestDistancesFromStart[current] < currentWeight {
			continue
		}

		destinations := bidirGraph[current]
		for next, weight := range destinations {
			nextWeight := currentWeight + weight
			if nextWeight < shortestDistancesFromStart[next] {
				shortestDistancesFromStart[next] = nextWeight
				heap.Push(&edgeQueue, Edge{nextWeight, next})
			}
		}
	}

	answer := make([]byte, 0, 12*N)
	for i := 2; i <= N; i++ {
		if i > 2 {
			answer = append(answer, ' ')
		}
		answer = strconv.AppendInt(answer, int64(shortestDistancesFromStart[i]), 10)
	}
	answer = append(answer, '\n')
	writer.Write(answer)
}
