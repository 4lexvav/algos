package main

import (
	"testing"
)

func Test(t *testing.T) {
	testCases := []struct {
		name     string
		matrix   [][]byte
		expected int
	}{
		{
			"1",
			[][]byte{{'1', '0', '1', '0', '0'}, {'1', '0', '1', '1', '1'}, {'1', '1', '1', '1', '1'}, {'1', '0', '0', '1', '0'}},
			6,
		},
		{
			"2",
			[][]byte{{'0'}},
			0,
		},
		{
			"3",
			[][]byte{{'1'}},
			1,
		},
		{
			"3",
			[][]byte{{'1', '1'}},
			2,
		},
		{
			"4",
			[][]byte{{'0', '0', '1', '0'}, {'0', '0', '1', '0'}, {'0', '0', '1', '0'}, {'0', '0', '1', '1'}, {'0', '1', '1', '1'}, {'0', '1', '1', '1'}, {'1', '1', '1', '1'}},
			9,
		},
		{
			"3",
			[][]byte{{'1', '1'}, {'1', '1'}},
			4,
		},
	}

	for _, test := range testCases {
		t.Run(test.name, func(t *testing.T) {
			result := maximalRectangle(test.matrix)
			if result != test.expected {
				t.Fatalf("Test case #%s failed: expected: %d got: %d", test.name, test.expected, result)
			}
		})
	}
}

func maximalRectangle(matrix [][]byte) int {
	if len(matrix) == 1 && len(matrix[0]) == 1 {
		return toInt(matrix[0][0])
	}

	maxArea := 0
	table := make([][][3]int, len(matrix))
	for i := range table {
		table[i] = make([][3]int, len(matrix[i]))
	}

	for i := range matrix {
		for j := range matrix[i] {
			if matrix[i][j] == '0' {
				table[i][j] = [3]int{0, 0, 0}
				continue
			}

			area := 0
			topX := 0
			topY := 0
			leftY := 0
			leftX := 0
			topArea := 0
			leftArea := 0
			raw := [3]int{}

			if i > 0 {
				topY = table[i-1][j][0]
				topX = table[i-1][j][1]
				topArea = table[i-1][j][2]
			}

			if j > 0 {
				leftY = table[i][j-1][0]
				leftX = table[i][j-1][1]
				leftArea = table[i][j-1][2]
			}

			if topArea > 0 {
				raw[0] = topY + 1
			}

			if leftArea > 0 {
				raw[1] = leftX + 1
			}

			table[i][j] = raw
			area = max(raw[0], raw[1]) + 1
			if leftY > 0 && topX > 0 {
				height := min(raw[0], leftY)
				length := min(raw[1], topX)
				colsCount, rowsCount := countLines(table, height, length, i, j)
				area = max(area, (height+1)*colsCount, (length+1)*rowsCount)
			}

			table[i][j][2] = area
			if area > maxArea {
				maxArea = area
			}
		}
	}

	return maxArea
}

func countLines(table [][][3]int, height, length, i, j int) (int, int) {
	colsCount := 0
	for n := j; n >= 0 && table[i][n][0] >= height; n-- {
		colsCount++
	}

	rowsCount := 0
	for n := i; n >= 0 && table[n][j][1] >= length; n-- {
		rowsCount++
	}

	return colsCount, rowsCount
}

func toInt(v byte) int {
	if v == '1' {
		return 1
	}
	return 0
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(nums ...int) int {
	max := nums[0]
	for i := range nums {
		if nums[i] > max {
			max = nums[i]
		}
	}
	return max
}
