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

type Queue struct {
	nodes []int
}

func (q *Queue) Push(v int) {
	q.nodes = append(q.nodes, v)
}
func (q *Queue) Pop() int {
	v := q.nodes[0]
	q.nodes = q.nodes[1:]
	return v
}

func readInt() int {
	scanner.Scan()
	res, _ := strconv.Atoi(scanner.Text())
	return res
}

func main() {
	scanner = bufio.NewScanner(os.Stdin)
	writer = bufio.NewWriter(os.Stdout)

	defer writer.Flush()

	const maxCapacity = 5 * 1024 * 1024
	buf := make([]byte, maxCapacity)
	scanner.Buffer(buf, maxCapacity)
	scanner.Split(bufio.ScanWords)

	type Set = map[int]bool

	T := readInt()
	for range T {
		N := readInt()
		M := readInt()

		edges := make([]Set, N+1)
		nodeColors := make([]int, N+1) // 0 == Not visited, color: 1 or -1

		for range M {
			a := readInt()
			b := readInt()
			if edges[a] == nil {
				edges[a] = Set{}
			}
			if edges[b] == nil {
				edges[b] = Set{}
			}
			edges[a][b] = false
			edges[b][a] = false
		}

		q := Queue{}
		q.Push(1)
		nodeColors[1] = 1
		for len(q.nodes) > 0 {
			currentNode := q.Pop()
			currentColor := nodeColors[currentNode]
			for dest := range edges[currentNode] {
				if nodeColors[dest] == 0 {
					nodeColors[dest] = currentColor * -1
					edges[currentNode][dest] = true
					edges[dest][currentNode] = true

					q.Push(dest)
				}
			}
		}

		var pair [2]int
	outer:
		for start, destinations := range edges[1:] {
			for dest, visited := range destinations {
				if !visited && nodeColors[start+1] == nodeColors[dest] {
					pair = [2]int{start + 1, dest}
					break outer
				}
			}
		}

		if pair[0] == 0 && pair[1] == 0 {
			fmt.Fprintln(writer, "-1")
			continue
		}

		start := pair[0]
		end := pair[1]
		parents := make([]int, N+1)
		q = Queue{}
		q.Push(start)
		nodeColors[start] = 2
		for len(q.nodes) > 0 {
			cur := q.Pop()
			if cur == end {
				break
			}
			for dest, visited := range edges[cur] {
				if visited && nodeColors[dest] != 2 {
					parents[dest] = cur
					nodeColors[dest] = 2
					q.Push(dest)
				}
			}
		}
		s := end
		n := 1
		answer := make([]byte, 0, 2*N)
		for {
			answer = strconv.AppendInt(answer, int64(s), 10)
			if parents[s] != 0 {
				s = parents[s]
				answer = append(answer, ' ')
				n++
			} else {
				break
			}
		}
		answer = append(answer, '\n')
		fmt.Fprintln(writer, n)
		writer.Write(answer)
	}
}
