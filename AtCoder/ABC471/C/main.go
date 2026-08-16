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

type Pos struct {
	Num   int
	Right int
	Left  int
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
	nums := make([]Pos, N+1)
	nums[0] = Pos{0, 0, 0}
	for i := range N {
		nums[i+1] = Pos{readInt(), 0, 0}
	}
	sort.Slice(nums, func(i, j int) bool {
		return nums[i].Num < nums[j].Num
	})

	curIdx := 0
	readCnt := 0
	dist := 0

	for i := range len(nums) {
		nums[i].Left = i - 1
		nums[i].Right = i + 1
		if nums[i].Num == 0 {
			curIdx = i
		}
	}

	for {
		//Start From 0 on {curIdx}
		curPos := nums[curIdx]
		leftIndex := curPos.Left
		rightIndex := curPos.Right
		if leftIndex < 0 {
			dist += nums[curPos.Right].Num - curPos.Num
			nums[curIdx].Right = nums[curPos.Right].Right
			nums[curPos.Right].Left = nums[curIdx].Left
			curIdx = curPos.Right

		} else if rightIndex > N {
			dist += curPos.Num - nums[curPos.Left].Num
			nums[curIdx].Left = nums[curPos.Left].Left
			nums[curPos.Left].Right = nums[curIdx].Right
			curIdx = curPos.Left

		} else {
			leftNum := nums[curPos.Left].Num
			rightNum := nums[curPos.Right].Num
			leftDist := curPos.Num - leftNum
			rightDist := rightNum - curPos.Num
			if leftDist <= rightDist {
				dist += curPos.Num - nums[curPos.Left].Num
				nums[curIdx].Left = nums[curPos.Left].Left
				nums[curPos.Left].Right = nums[curIdx].Right
				nums[curPos.Right].Left = nums[curIdx].Left
				curIdx = curPos.Left

			} else {
				dist += nums[curPos.Right].Num - curPos.Num
				nums[curIdx].Right = nums[curPos.Right].Right
				nums[curPos.Right].Left = nums[curIdx].Left
				nums[curPos.Left].Right = nums[curIdx].Right
				curIdx = curPos.Right
			}
		}
		readCnt++
		if readCnt == N {
			break
		}
	}
	fmt.Fprintln(writer, dist)
}
