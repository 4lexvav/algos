package main

import (
	"testing"
)

func Test(t *testing.T) {
	testCases := []struct {
		name     string
		n        int
		expected int
	}{
		{
			"1",
			2,
			2,
		},
		{
			"2",
			3,
			3,
		},
	}

	for _, test := range testCases {
		t.Run(test.name, func(t *testing.T) {
			result := climbStairs(test.n)
			if result != test.expected {
				t.Fatalf("Test case #%s failed: expected: %d got: %d", test.name, test.expected, result)
			}
		})
	}
}

func climbStairs(n int) int {
	a, b := 1, 1
	for ; n > 0; n-- {
		a, b = b, a+b
	}
	return a
}

func climbStairsWithStore(n int) int {
	if n <= 3 {
		return n
	}

	ways := make([]int, n)
	ways[0] = 1
	ways[1] = 2
	for i := 2; i < n; i++ {
		ways[i] = ways[i-1] + ways[i-2]
	}

	return ways[n-1]
}
