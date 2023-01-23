package main

import (
	"testing"
)

func Test(t *testing.T) {
	testCases := []struct {
		name     string
		m        int
		n        int
		expected int
	}{
		{
			"1",
			3,
			7,
			28,
		},
		{
			"2",
			3,
			2,
			3,
		},
		{
			"3",
			1,
			7,
			1,
		},
	}

	for _, test := range testCases {
		t.Run(test.name, func(t *testing.T) {
			//result := uniquePaths(test.m, test.n)
			result := uniquePathsOptimized(test.m, test.n)
			if result != test.expected {
				t.Fatalf("Test case #%s failed: expected: %d got: %d", test.name, test.expected, result)
			}
		})
	}
}

func uniquePaths(m int, n int) int {
	matrix := matrix(m, n)
	for i := 0; i < m; i++ {
		matrix[i][0] = 1
	}

	for i := 0; i < n; i++ {
		matrix[0][i] = 1
	}

	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			matrix[i][j] = matrix[i][j-1] + matrix[i-1][j]
		}
	}

	return matrix[m-1][n-1]
}

func uniquePathsOptimized(m int, n int) int {
	row := make([]int, n)
	for i := 0; i < n; i++ {
		row[i] = 1
	}

	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			row[j] = row[j-1] + row[j]
		}
	}

	return row[n-1]
}

func matrix(m, n int) [][]int {
	matrix := make([][]int, m)
	for i := 0; i < m; i++ {
		matrix[i] = make([]int, n)
	}
	return matrix
}
