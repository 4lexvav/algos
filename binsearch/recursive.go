package main

import (
	"fmt"
)

func main() {
	fmt.Println(binSearchRec([]int{1, 2, 3, 4, 8, 9, 11, 14, 18, 23, 45, 57, 79, 88}, 4))
}

func binSearchRec(nums []int, n int) int {
	if len(nums) == 0 {
		return -1
	}
	m := len(nums) / 2
	if nums[m] == n {
		return m
	}

	if nums[m] > n {
		return binSearchRec(nums[0:m], n)
	}
	pos := binSearchRec(nums[m+1:], n)
	if pos >= 0 {
		pos += m + 1
	}
	return pos
}
