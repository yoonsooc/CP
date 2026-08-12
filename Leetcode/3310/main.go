package main

import (
	"fmt"
)

func bfs(graph map[int][]int, visited []bool, inDegree []int, start int) {
	n := len(graph)
	q := make([]int, 0, n)
	head := 0

	visited[start] = true
	q = append(q, start)

	for head < len(q) {
		cur := q[head]
		head++

		for _, next := range graph[cur] {
			inDegree[next]--
			if visited[next] {
				continue
			}
			visited[next] = true
			q = append(q, next)
		}
	}
}

func remainingMethods(n int, k int, invocations [][]int) []int {
	graph := make(map[int][]int, n)
	inDegree := make([]int, n)

	for _, edge := range invocations {
		f := edge[0]
		t := edge[1]

		graph[f] = append(graph[f], t)
		inDegree[t]++
	}

	visited := make([]bool, n)
	bfs(graph, visited, inDegree, k)

	removable := true
	result := []int{}
	for node, v := range visited {
		if v && inDegree[node] > 0 {
			removable = false
			break
		} else if !v {
			result = append(result, node)
		}
	}
	if !removable {
		allNodes := make([]int, n)
		for i := range n {
			allNodes[i] = i
		}
		return allNodes
	}
	return result
}

func main() {
	fmt.Println(remainingMethods(3, 2, [][]int{{1, 0}, {2, 0}}))
	fmt.Println(remainingMethods(4, 1, [][]int{{1, 2}, {0, 1}, {3, 2}}))
	fmt.Println(remainingMethods(5, 0, [][]int{{1, 2}, {0, 2}, {0, 1}, {3, 4}}))
	fmt.Println(remainingMethods(3, 2, [][]int{{1, 2}, {0, 1}, {2, 0}}))
}
