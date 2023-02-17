package main

import "testing"

func TestProblem(t *testing.T) {
	tests := []struct {
		name     string
		expected bool
	}{
		{
			"1",
			true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := problem()
			if result != tt.expected {
				t.Fatalf("Test case #%s failed, expected %v, got: %v", tt.name, tt.expected, result)
			}
		})
	}
}

func problem() bool {
	return true
}
