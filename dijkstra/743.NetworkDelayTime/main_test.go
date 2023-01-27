package main

import (
	"container/heap"
	"math"
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

func networkDelayTime(times [][]int, n int, k int) int {
	// 0 - target node index
	// 1 - cost of moving to this node
	graph := make([][][2]int, n)
	for i := range times {
		graph[times[i][0]-1] = append(graph[times[i][0]-1], [2]int{times[i][1] - 1, times[i][2]})
	}

	k--
	visited := make(map[int]struct{})
	costs := make([]int, n)
	for i := 0; i < n; i++ {
		costs[i] = math.MaxInt
	}
	costs[k] = 0

	h := graphHeap{}
	h.Push([2]int{k, 0})

	for h.Len() > 0 {
		from := heap.Pop(&h).([2]int)
		visited[from[0]] = struct{}{}

		for i := range graph[from[0]] {
			to := graph[from[0]][i]
			cost := costs[from[0]] + to[1]

			if cost < costs[to[0]] {
				costs[to[0]] = cost
			}

			if _, ok := visited[to[0]]; !ok {
				heap.Push(&h, to)
			}
		}
	}

	if len(visited) < n {
		return -1
	}

	return max(costs)
}

func max(costs []int) int {
	max := 0
	for i := range costs {
		if costs[i] > max {
			max = costs[i]
		}
	}

	return max
}

// [][0] - node index
// [][1] - cost of moving to this node
type graphHeap [][2]int

func (h graphHeap) Len() int {
	return len(h)
}

func (h graphHeap) Less(i, j int) bool {
	return h[i][1] < h[j][1]
}

func (h graphHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *graphHeap) Push(x interface{}) {
	*h = append(*h, x.([2]int))
}

func (h *graphHeap) Pop() interface{} {
	item := (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]
	return item
}
