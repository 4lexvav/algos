package main

import "testing"

func TestProblem(t *testing.T) {
	tests := []struct {
		name     string
		height   []int
		expected int
	}{
		{
			"1",
			[]int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1},
			6,
		},
		{
			"2",
			[]int{4, 2, 0, 3, 2, 5},
			9,
		},
		{
			"3",
			[]int{0, 5, 4, 3, 2, 1, 0},
			0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := trap(tt.height)
			if result != tt.expected {
				t.Fatalf("Test case #%s failed, expected %v, got: %v", tt.name, tt.expected, result)
			}
		})
	}
}

func trap(height []int) int {
	count, maxLeft, maxRight := 0, 0, 0
	for len(height) > 1 {
		if height[0] <= height[len(height)-1] {
			if height[0] > maxLeft {
				maxLeft = height[0]
			} else {
				count += maxLeft - height[0]
			}
			height = height[1:]
		} else {
			if height[len(height)-1] > maxRight {
				maxRight = height[len(height)-1]
			} else {
				count += maxRight - height[len(height)-1]
			}
			height = height[:len(height)-1]
		}
	}
	return count
}

func trapLeftRight(height []int) int {
	count := 0
	left, maxLeft := 0, 0
	right, maxRight := len(height)-1, 0

	for left < right {
		if height[left] <= height[right] {
			if height[left] > maxLeft {
				maxLeft = height[left]
			} else {
				count += maxLeft - height[left]
			}
			left++
		} else {
			if height[right] > maxRight {
				maxRight = height[right]
			} else {
				count += maxRight - height[right]
			}
			right--
		}
	}

	return count
}

func trapSlow(height []int) int {
	if len(height) < 3 {
		return 0
	}

	size := len(height)
	waterAmount := 0

	for i := 0; i < size-1; i++ {
		if !isValid(height[i], height[i+1]) {
			continue
		}

		leftBarIdx := i
		rightBarIdx := findRightBar(height[i], height[i+1:])
		if rightBarIdx == 0 {
			continue
		}

		// we should add i + 1 to rightBar because rightIndex is shrinked
		rightBarIdx += i + 1

		threshold := min(height[leftBarIdx], height[rightBarIdx])
		waterAmount += fillWater(height[leftBarIdx+1:rightBarIdx], threshold)
		i = rightBarIdx - 1
	}

	return waterAmount
}

func fillWater(heights []int, threshold int) int {
	filledCount := 0
	for i := range heights {
		count := threshold - heights[i]
		if count > 0 {
			filledCount += count
		}
	}

	return filledCount
}

func findRightBar(leftBar int, heights []int) int {
	if len(heights) == 0 {
		return 0
	}

	size := len(heights)
	rightBarIdx := 0
	maxHeight := 0

	for i := 0; i < size; i++ {
		if heights[i] > maxHeight {
			rightBarIdx = i
			maxHeight = heights[i]
		}

		if heights[i] >= leftBar {
			break
		}
	}

	return rightBarIdx
}

func isValid(leftHeight, rightHeight int) bool {
	return leftHeight > 0 && leftHeight > rightHeight
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}
