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

// U, L, D, R
var rIdx = [4]int{-1, 0, 1, 0}
var cIdx = [4]int{0, -1, 0, 1}

type Footprint struct {
	visited bool
	prevDir int // 0-3
}
type Direction struct {
	i    int
	j    int
	from int // 0~3
}

type Char rune

func (c Char) String() string {
	return string(rune(c))
}

type Grid[T fmt.Stringer] struct {
	data [][]T
}

func (g *Grid[T]) Print() {
	for _, row := range g.data {
		for _, cell := range row {
			fmt.Printf("%s", cell.String())
		}
		fmt.Println()
	}
}

type Queue []Direction

func (q *Queue) Push(d Direction) {
	*q = append(*q, d)
}
func (q *Queue) Pop() Direction {
	d := (*q)[0]
	*q = (*q)[1:]
	return d
}

func isIn(i, j, H, W int) bool {
	return i >= 0 && i < H && j >= 0 && j < W
}

func main() {
	scanner = bufio.NewScanner(os.Stdin)
	writer = bufio.NewWriter(os.Stdout)

	defer writer.Flush()

	const maxCapacity = 5 * 1024 * 1024
	buf := make([]byte, maxCapacity)
	scanner.Buffer(buf, maxCapacity)
	scanner.Split(bufio.ScanWords)

	H := readInt()
	W := readInt()

	var start [2]int
	goal := [2]int{-1, -1}
	visited := make([][][]Footprint, H)
	data := make([][]Char, H)
	for i := range H {
		scanner.Scan()
		line := scanner.Text()
		visited[i] = make([][]Footprint, W)
		data[i] = make([]Char, W)
		for j, v := range line {
			data[i][j] = Char(v)
			if v == 'S' {
				start = [2]int{i, j}
			}
			visited[i][j] = []Footprint{
				{false, 0},
				{false, 0},
				{false, 0},
				{false, 0},
			}
		}
	}

	q := Queue{}
	for d := range 4 {
		ni := start[0] + rIdx[d]
		nj := start[1] + cIdx[d]
		if isIn(ni, nj, H, W) {
			q.Push(Direction{ni, nj, d})
			visited[ni][nj][d] = Footprint{true, -1}
		}
	}

	goalFrom := 0
outer:
	for len(q) > 0 {
		curLen := len(q)
		for range curLen {
			curNode := q.Pop()
			i := curNode.i
			j := curNode.j
			from := curNode.from
			curVal := data[i][j]
			switch curVal {
			case 'G':
				goal = [2]int{i, j}
				goalFrom = from
				break outer
			case 'S', '.':
				for d := range 4 {
					ni := rIdx[d] + i
					nj := cIdx[d] + j
					if isIn(ni, nj, H, W) {
						if data[ni][nj] != '#' && visited[ni][nj][d].visited == false {
							q.Push(Direction{ni, nj, d})
							visited[ni][nj][d] = Footprint{true, from}
						}
					}
				}
			case 'o':
				ni := rIdx[from] + i
				nj := cIdx[from] + j
				if isIn(ni, nj, H, W) {
					if data[ni][nj] != '#' && visited[ni][nj][from].visited == false {
						q.Push(Direction{ni, nj, from})
						visited[ni][nj][from] = Footprint{true, from}
					}
				}
			case 'x':
				for d := range 4 {
					if d == from {
						continue
					}
					ni := rIdx[d] + i
					nj := cIdx[d] + j
					if isIn(ni, nj, H, W) {
						if data[ni][nj] != '#' && visited[ni][nj][d].visited == false {
							q.Push(Direction{ni, nj, d})
							visited[ni][nj][d] = Footprint{true, from}
						}
					}
				}
			}
		}
	}

	if goal[0] == -1 {
		fmt.Fprint(writer, "No")
		return
	}

	// fmt.Println(visited)
	// Find Path
	dirChar := [4]byte{'U', 'L', 'D', 'R'}
	path := []byte{}
	i, j, d := goal[0], goal[1], goalFrom
	for d != -1 {
		path = append(path, dirChar[d])
		pi, pj := i-rIdx[d], j-cIdx[d]
		d = visited[i][j][d].prevDir
		i, j = pi, pj
	}
	for l, r := 0, len(path)-1; l < r; l, r = l+1, r-1 {
		path[l], path[r] = path[r], path[l]
	}
	fmt.Fprintln(writer, "Yes")
	fmt.Fprintln(writer, string(path))
}
