package main

import (
	"fmt"
)

func maxSubarrayLength(nums []int, k int) int {
	answer, left := 0, 0
	elemCount := make(map[int]int, len(nums))

	for right := range len(nums) {
		current := nums[right]
		elemCount[current]++
		for elemCount[current] > k {
			leftEnd := nums[left]
			elemCount[leftEnd]--
			if elemCount[leftEnd] == 0 {
				delete(elemCount, leftEnd)
			}
			left++
		}
		answer = max(answer, right-left+1)
	}
	return answer
}

func main() {
	fmt.Println(maxSubarrayLength([]int{1, 4, 4, 3}, 1))             // Output: 4
	fmt.Println(maxSubarrayLength([]int{1, 2, 3, 1, 2, 3, 1, 2}, 2)) // Output: 6
	fmt.Println(maxSubarrayLength([]int{1, 2, 1, 2, 1, 2, 1, 2}, 1)) // Output: 2
	fmt.Println(maxSubarrayLength([]int{5, 5, 5, 5, 5, 5, 5}, 4))    // Output: 4
}
