// You can edit this code!
// Click here and start typing.
package main

import (
	"fmt"
	"math/rand"
)

func main() {
	nums := []int{5, 3, 4, 1, 8, 9, 2, 11}
	fmt.Println(nums)
	fmt.Println(quickSort(nums))
	fmt.Println()

	nums = []int{5, 3, 4, 1, 8, 9, 2, 11}
	fmt.Println(nums)
	fmt.Println(quickSortInplace(nums))
	fmt.Println(nums)
}

func quickSort(nums []int) []int {
	if len(nums) < 2 {
		return nums
	}

	//pivot := rand.Int() % len(nums)
	pivot := 0
	left := make([]int, 0)
	right := make([]int, 0)

	//nums = append(nums[:pivot], nums[pivot+1:]...)

	for _, v := range nums[1:] {
		if v <= nums[pivot] {
			left = append(left, v)
		} else {
			right = append(right, v)
		}
	}

	left = append(quickSort(left), nums[pivot])
	return append(left, quickSort(right)...)
}

func quickSortInplace(nums []int) []int {
	if len(nums) < 2 {
		return nums
	}

	pivot := rand.Int() % len(nums)
	left, right := 0, len(nums)-1

	// shift pivot to the end
	nums[pivot], nums[right] = nums[right], nums[pivot]

	// fill left part
	for i := range nums {
		if nums[i] < nums[right] {
			nums[left], nums[i] = nums[i], nums[left]
			left++
		}
	}

	// place pivot between left and right
	nums[left], nums[right] = nums[right], nums[left]

	quickSortInplace(nums[:left])
	quickSortInplace(nums[left+1:])

	return nums
}
