package main

import (
	"testing"
)

func Test(t *testing.T) {
	testCases := []struct {
		name     string
		times    [][]int
		size     int
		src      int
		expected int
	}{
		{
			"1",
			[][]int{{2, 1, 1}, {2, 3, 1}, {3, 4, 1}},
			4,
			2,
			2,
		},
		{
			"2",
			[][]int{{1, 2, 1}},
			2,
			1,
			1,
		},
		{
			"3",
			[][]int{{1, 2, 1}},
			2,
			2,
			-1,
		},
		{
			"3",
			[][]int{{1, 2, 1}, {2, 1, 3}},
			2,
			2,
			3,
		},
	}

	for _, test := range testCases {
		t.Run(test.name, func(t *testing.T) {
			result := networkDelayTime(test.times, test.size, test.src)
			if result != test.expected {
				t.Fatalf("Test case #%s failed: expected: %d got: %d", test.name, test.expected, result)
			}
		})
	}
}

func networkDelayTime(times [][]int, size, src int) int {
	return 0
}
