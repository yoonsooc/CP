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
	K := readInt()

	grid := make([][]int, H)
	partSum := make([][]int, H)
	for i := range H {
		grid[i] = make([]int, W)
		partSum[i] = make([]int, W)

		scanner.Scan()
		line := scanner.Text()
		for j, s := range line {
			n, _ := strconv.Atoi(string(s))
			grid[i][j] = n
			if i > 0 {
				partSum[i][j] += partSum[i-1][j]
			}
			if j > 0 {
				partSum[i][j] += partSum[i][j-1]
			}
			if i > 0 && j > 0 {
				partSum[i][j] -= partSum[i-1][j-1]
			}
			if n == 1 {
				partSum[i][j] += 1
			}
		}
	}

	// fmt.Println(K, grid)
	// fmt.Println(partSum)
	count := 0
	freq := make([]int, H*W+1)
	freqExist := make([]int, W)
	for lb := range H {
		for ub := lb; ub < H; ub++ {
			// 이번 행 구간에서 최초 열부터 돌면서 구간 합을 저장,
			//  freq := make(map[int]int)
			// [BAD] for loop의 총 cell 방문 횟수 > 2*10^8.
			// hash calculation, Bucket search가 일어나는 map 연산(> 10ns)을 cell 마다 2회 이상 수행하기만 해도 4s 초과

			// 누적합 한도만큼의 크기를 갖는 슬라이스를 미리 할당 후 인덱스로 접근, 매 행 구간마다 초기화가 아닌 +/-연산으로 값 복원
			// freq := make([]int, H*W+1)
			freq[0]++
			freqExist = append(freqExist[:0], 0)

			for col := range W {
				currentSum := partSum[ub][col]
				if lb > 0 {
					currentSum -= partSum[lb-1][col]
				}

				target := currentSum - K
				if target >= 0 {
					count += freq[target]
				}
				freq[currentSum] = freq[currentSum] + 1
				freqExist = append(freqExist, currentSum)
			}

			for _, s := range freqExist {
				freq[s]--
			}
		}
	}
	fmt.Fprintln(writer, count)
}
