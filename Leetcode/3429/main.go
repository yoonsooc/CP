package main

import (
	"fmt"
	"math"
)

func makeMatrix(rows, cols int) [][]int64 {
	matrix := make([][]int64, rows)
	for i := range rows {
		matrix[i] = make([]int64, cols)
	}
	return matrix
}

func getMinElem(arr []int64) int64 {
	if len(arr) == 0 {
		panic("array is empty")
	}

	var min int64 = math.MaxInt64
	for _, v := range arr {
		if v < min {
			min = v
		}
	}
	return min
}

func minCost(n int, cost [][]int) int64 {
	var length = n / 2
	// There are 6 possible combinations of colors for two adjacent houses:
	var cases = [6][2]int{
		{0, 1},
		{0, 2},
		{1, 0},
		{1, 2},
		{2, 0},
		{2, 1},
	}
	const caseCount = len(cases)
	const colorCount = 3

	var memo = makeMatrix(length, caseCount)
	for i := range length {
		for j := range caseCount {
			memo[i][j] = math.MaxInt64
		}
	}

	for i := range length {
		// ->			<-
		// [0] 1 2 3 4 [5], length = 3
		var leftFromLine = i
		var rightFromLine = length*2 - i - 1

		for j, colorCase := range cases {
			leftColor := colorCase[0]
			rightColor := colorCase[1]

			memo[i][j] = int64(cost[leftFromLine][leftColor]) + int64(cost[rightFromLine][rightColor])
			// fmt.Printf("on (%d, %d): color(%d, %d): %d\n", leftFromLine, rightFromLine, leftColor, rightColor, memo[i][j])
			if i == 0 {
				continue
			}

			var minimumPrev int64 = math.MaxInt64
			for k, prevColorCase := range cases {
				prevLeftColor := prevColorCase[0]
				prevRightColor := prevColorCase[1]

				if leftColor == prevLeftColor || rightColor == prevRightColor {
					continue
				}

				// fmt.Printf("compare with (%d, %d): %d\n", prevLeftColor, prevRightColor, memo[i-1][k])
				if memo[i-1][k] < minimumPrev {
					minimumPrev = memo[i-1][k]
				}
			}
			memo[i][j] += minimumPrev
		}
	}

	// fmt.Println("memo:", memo)
	return getMinElem(memo[length-1])
}

func main() {
	fmt.Println(minCost(4, [][]int{{3, 5, 7}, {6, 2, 9}, {4, 8, 1}, {7, 3, 5}}))
	fmt.Println(minCost(6, [][]int{{2, 4, 6}, {5, 3, 8}, {7, 1, 9}, {4, 6, 2}, {3, 5, 7}, {8, 2, 4}}))
}
