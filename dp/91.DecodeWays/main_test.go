package main

import (
	"strconv"
	"testing"
)

func Test(t *testing.T) {
	testCases := []struct {
		name     string
		s        string
		expected int
	}{
		{
			"1",
			"226",
			3,
		},
		{
			"2",
			"12",
			2,
		},
		{
			"3",
			"06",
			0,
		},
		{
			"4",
			"20419",
			2,
		},
		{
			"5",
			"9012",
			0,
		},
		{
			"6",
			"27",
			1,
		},
	}

	for _, test := range testCases {
		t.Run(test.name, func(t *testing.T) {
			result := numDecodings2(test.s)
			if result != test.expected {
				t.Fatalf("Test case #%s failed: expected: %d got: %d", test.name, test.expected, result)
			}
		})
	}
}

func numDecodings(s string) int {
	if s[0] == '0' {
		return 0
	}

	n := len(s)
	count1 := 1
	count2 := 0
	for i := 1; i < n; i++ {
		prevCount1 := count1
		if s[i] == '0' {
			count1 = 0
		} else {
			count1 += count2
		}

		// reset if first num is 0 or greater than 2 or whole num is greater than 26
		if s[i-1] < '1' || s[i-1] > '2' || (s[i-1] == '2' && s[i] > '6') {
			count2 = 0
		} else {
			count2 = prevCount1
		}
	}

	return count1 + count2
}

func numDecodings2(s string) int {
	dp := make([]int, len(s)+1)
	dp[0] = 1
	if s[0] != '0' {
		dp[1] = 1
	}

	for i := 2; i < len(dp); i++ {
		a, _ := strconv.Atoi(s[i-1 : i])
		b, _ := strconv.Atoi(s[i-2 : i])
		if a >= 1 {
			dp[i] += dp[i-1]
		}
		if b >= 10 && b <= 26 {
			dp[i] += dp[i-2]
		}
	}

	return dp[len(s)]
}
