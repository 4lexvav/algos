package main

import (
	"testing"
)

func TestReverse(t *testing.T) {
	testCases := []struct {
		name     string
		str      string
		expected string
	}{
		{
			"1",
			"name",
			"eman",
		},
		{
			"2",
			"foo",
			"oof",
		},
		{
			"3",
			"foobar",
			"raboof",
		},
		{
			"4",
			"advertisement",
			"tnemesitrevda",
		},
		{
			"5",
			"The quick brown 狐 jumped over the lazy 犬",
			"犬 yzal eht revo depmuj 狐 nworb kciuq ehT",
		},
	}

	for _, test := range testCases {
		t.Run(test.name, func(t *testing.T) {
			result := reverse(test.str)
			if result != test.expected {
				t.Fatalf("Test case #%s failed: expected: %s got: %s", test.name, test.expected, result)
			}
		})
	}
}

func reverse(str string) string {
	chars := []rune(str)
	for i, j := 0, len(chars)-1; i < j; i, j = i+1, j-1 {
		chars[i], chars[j] = chars[j], chars[i]
	}

	return string(chars)
}
