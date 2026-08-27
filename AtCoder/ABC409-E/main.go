package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
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

type Queue []int

func (q *Queue) Push(d int) {
	*q = append(*q, d)
}
func (q *Queue) Pop() int {
	d := (*q)[0]
	*q = (*q)[1:]
	return d
}

type Edge struct {
	vert   int
	weight int
}

type Node struct {
	vert             int
	parent           int
	parentEdgeWeight int
}

func abs(i int) int {
	if i < 0 {
		return i * (-1)
	}
	return i
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
	trons := make([]int, N)
	for i := range N {
		trons[i] = readInt()
	}
	edges := make([][]Edge, N)
	for range N - 1 {
		u := readInt() - 1
		v := readInt() - 1
		w := readInt()
		if edges[u] == nil {
			edges[u] = make([]Edge, 0)
		}
		if edges[v] == nil {
			edges[v] = make([]Edge, 0)
		}
		edges[u] = append(edges[u], Edge{v, w})
		edges[v] = append(edges[v], Edge{u, w})
	}

	// BFS From 0(Random Root) & Find Parent Node
	start := 0
	q := Queue{}
	q.Push(start)
	order := make([]int, 0)
	visited := make([]bool, N)
	visited[start] = true
	tree := make([]Node, N)
	tree[start] = Node{0, -1, 0}
	for len(q) > 0 {
		u := q.Pop()
		order = append(order, u)
		for _, edge := range edges[u] {
			v := edge.vert
			if visited[v] == false {
				tree[v] = Node{v, u, edge.weight}
				visited[v] = true
				q.Push(v)
			}
		}
	}

	// Calculate Annihilation From Leaves
	cost := 0
	for v := range slices.Backward(order) {
		u := order[v]
		currentTrons := trons[u]
		par := tree[u].parent
		if par != -1 {
			trons[par] += currentTrons
			additionalCost := abs(currentTrons) * tree[u].parentEdgeWeight
			cost += additionalCost
		}
	}

	fmt.Println(cost)
}
