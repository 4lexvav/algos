package main

import (
	"reflect"
	"testing"
)

func TestDecrypt(t *testing.T) {
	testCases := []struct {
		name     string
		code     []int
		k        int
		expected []int
	}{
		{
			"1",
			[]int{5, 7, 1, 4},
			3,
			[]int{12, 10, 16, 13},
		},
		{
			"2",
			[]int{1, 2, 3, 4},
			0,
			[]int{0, 0, 0, 0},
		},
		{
			"3",
			[]int{2, 4, 9, 3},
			-2,
			[]int{12, 5, 6, 13},
		},
	}

	for _, test := range testCases {
		t.Run(test.name, func(t *testing.T) {
			result := decrypt(test.code, test.k)
			if !reflect.DeepEqual(result, test.expected) {
				t.Fatalf("Test case #%s failed: expected: %d got: %d", test.name, test.expected, result)
			}
		})
	}
}

func decrypt(code []int, k int) []int {
	n := len(code)

	if k == 0 {
		return make([]int, n)
	}

	absK := abs(k)
	result := make([]int, n)

	for i := range code {
		from := i + 1
		if k < 0 {
			from = i - absK + n
		}

		for j := 0; j < absK; j++ {
			result[i] += code[(from+j)%n]
		}
	}

	return result
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
