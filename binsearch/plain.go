package main

import "fmt"

func main() {
	fmt.Println(binSearch([]int{1, 2, 3, 4, 8, 9, 11, 14, 18, 23, 45, 57, 79, 88}, 88))
}

func binSearch(nums []int, n int) int {
	l := 0
	r := len(nums) - 1

	for l <= r {
		m := (l + r) / 2
		if nums[m] == n {
			return m
		} else if nums[m] > n {
			r = m - 1
		} else {
			l = m + 1
		}
	}

	return -1
}
