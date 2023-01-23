package main

import (
	"container/heap"
	"testing"
)

func Test(t *testing.T) {
	testCases := []struct {
		name     string
		heights  [][]int
		expected int
	}{
		{
			"1",
			[][]int{{1, 2, 2}, {3, 8, 2}, {5, 3, 5}},
			2,
		},
		{
			"2",
			[][]int{{1, 2, 3}, {3, 8, 4}, {5, 3, 5}},
			1,
		},
		{
			"3",
			[][]int{{1, 2, 1, 1, 1}, {1, 2, 1, 2, 1}, {1, 2, 1, 2, 1}, {1, 2, 1, 2, 1}, {1, 1, 1, 2, 1}},
			0,
		},
		{
			"4",
			[][]int{{1, 10, 6, 7, 9, 10, 4, 9}},
			9,
		},
		{
			"5",
			[][]int{{4, 3, 4, 10, 5, 5, 9, 2}, {10, 8, 2, 10, 9, 7, 5, 6}, {5, 8, 10, 10, 10, 7, 4, 2}, {5, 1, 3, 1, 1, 3, 1, 9}, {6, 4, 10, 6, 10, 9, 4, 6}},
			5,
		},
	}

	for _, test := range testCases {
		t.Run(test.name, func(t *testing.T) {
			result := minimumEffortPath(test.heights)
			if result != test.expected {
				t.Fatalf("Test case #%s failed: expected: %d got: %d", test.name, test.expected, result)
			}
		})
	}
}

func minimumEffortPath(heights [][]int) int {
	rowsLen := len(heights)
	colsLen := len(heights[0])
	size := rowsLen * colsLen
	endCoord := size - 1
	visited := make([]bool, size)
	efforts := make([]int, size)

	hh := heightsHeap{}
	hh.Push([]int{0, 0, 0})

	for hh.Len() > 0 {
		height := heap.Pop(&hh).([]int)
		y := height[0]
		x := height[1]
		coord := y*colsLen + x
		if coord == endCoord {
			break
		}

		// mark cell visited
		visited[coord] = true

		for _, cell := range [4][2]int{{y - 1, x}, {y, x - 1}, {y + 1, x}, {y, x + 1}} {
			// skip if cell coordinates wrong or cell already visited
			if cell[0] < 0 || cell[1] < 0 || cell[0] >= rowsLen || cell[1] >= colsLen ||
				visited[cell[0]*colsLen+cell[1]] {
				continue
			}

			cellEffort := abs(heights[cell[0]][cell[1]] - heights[y][x])
			cellCoord := cell[0]*colsLen + cell[1]

			// skip this cell if it already has fewer efforts from another path
			if efforts[cellCoord] > 0 && efforts[cellCoord] < cellEffort {
				continue
			}

			// always choose effort with max value
			if cellEffort < efforts[coord] {
				cellEffort = efforts[coord]
			}

			efforts[cellCoord] = cellEffort
			heap.Push(&hh, []int{cell[0], cell[1], efforts[cellCoord]})
		}
	}

	return efforts[endCoord]
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

// [][0] row idx
// [][1] col idx
// [][2] effort needed to move to this height
type heightsHeap [][]int

func (h heightsHeap) Len() int {
	return len(h)
}

func (h heightsHeap) Less(i, j int) bool {
	return h[i][2] < h[j][2]
}

func (h heightsHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *heightsHeap) Push(x interface{}) {
	*h = append(*h, x.([]int))
}

func (h *heightsHeap) Pop() interface{} {
	col := (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]
	return col
}
