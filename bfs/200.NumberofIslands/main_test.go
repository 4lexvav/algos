package main

import "testing"

func TestNumberOfIslands(t *testing.T) {
	tests := []struct {
		name     string
		input    [][]byte
		expected int
	}{
		{
			"1",
			[][]byte{{'1', '1', '1', '1', '0'}, {'1', '1', '0', '1', '0'}, {'1', '1', '0', '0', '0'}, {'0', '0', '0', '0', '0'}},
			1,
		},
		{
			"2",
			[][]byte{{'1', '1', '0', '0', '0'}, {'1', '1', '0', '0', '0'}, {'0', '0', '1', '0', '0'}, {'0', '0', '0', '1', '1'}},
			3,
		},
		{
			"3",
			[][]byte{{'1', '1', '1'}, {'0', '1', '0'}, {'1', '1', '1'}},
			1,
		},
		{
			"4",
			[][]byte{{'1', '0', '1', '1', '1'}, {'1', '0', '1', '0', '1'}, {'1', '1', '1', '0', '1'}},
			1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if result := numIslands(tt.input); result != tt.expected {
				t.Fatalf("Test case #%s failed: expected: %d got: %d", tt.name, tt.expected, result)
			}
		})
	}
}

func numIslands(grid [][]byte) int {
	numIslands := 0
	m, n := len(grid), len(grid[0])
	// x/y coordinate of next cell containing one
	stack := make([][]int, 0, m*n)

	for i := range grid {
		for j := range grid[i] {
			if grid[i][j] == '0' {
				continue
			}

			numIslands++
			stack = append(stack, []int{i, j})

			for len(stack) > 0 {
				coord := stack[len(stack)-1]
				stack = stack[:len(stack)-1]
				x, y := coord[0], coord[1]
				grid[x][y] = '0'

				leftIdx := y - 1
				rightIdx := y + 1
				topIdx := x - 1
				bottomIdx := x + 1

				if leftIdx >= 0 && grid[x][leftIdx] == '1' {
					stack = append(stack, []int{x, leftIdx})
				}

				if rightIdx < n && grid[x][rightIdx] == '1' {
					stack = append(stack, []int{x, rightIdx})
				}

				if topIdx >= 0 && grid[topIdx][y] == '1' {
					stack = append(stack, []int{topIdx, y})
				}

				if bottomIdx < m && grid[bottomIdx][y] == '1' {
					stack = append(stack, []int{bottomIdx, y})
				}
			}
		}
	}

	return numIslands
}

func numIslandsWithVisited(grid [][]byte) int {
	numIslands := 0
	m, n := len(grid), len(grid[0])
	// visited[n] corresponds grid cell number
	visited := make(map[int]bool, m*n)
	// x/y coordinate of next cell containing one
	stack := make([][]int, 0, m*n)

	for i := range grid {
		for j := range grid[i] {
			//cell ID formula
			cellId := i*n + j
			if visited[cellId] {
				continue
			}

			if grid[i][j] == '1' {
				numIslands++
				stack = append(stack, []int{i, j})
			}

			for len(stack) > 0 {
				coord := stack[len(stack)-1]
				stack = stack[:len(stack)-1]
				x, y := coord[0], coord[1]
				cellId := x*n + y
				if visited[cellId] {
					continue
				}

				visited[cellId] = true

				leftIdx := y - 1
				rightIdx := y + 1
				topIdx := x - 1
				bottomIdx := x + 1

				if leftIdx >= 0 && grid[x][leftIdx] == '1' {
					stack = append(stack, []int{x, leftIdx})
				}

				if rightIdx < n && grid[x][rightIdx] == '1' {
					stack = append(stack, []int{x, rightIdx})
				}

				if topIdx >= 0 && grid[topIdx][y] == '1' {
					stack = append(stack, []int{topIdx, y})
				}

				if bottomIdx < m && grid[bottomIdx][y] == '1' {
					stack = append(stack, []int{bottomIdx, y})
				}
			}
		}
	}

	return numIslands
}

func numIslandsWrong(grid [][]byte) int {
	numIslands := 0

	for i := range grid {
		for j := range grid[i] {
			if grid[i][j] == '0' {
				continue
			}

			isTopZero := (i == 0) || grid[i-1][j] == '0'
			isLeftZero := (j == 0) || grid[i][j-1] == '0'

			if isTopZero && isLeftZero {
				numIslands++
			}
		}
	}

	return numIslands
}
