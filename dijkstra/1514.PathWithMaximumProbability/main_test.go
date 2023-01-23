package main

import (
	"container/heap"
	"testing"
)

func Test(t *testing.T) {
	testCases := []struct {
		name     string
		n        int
		edges    [][]int
		succProb []float64
		start    int
		end      int
		expected float64
	}{
		{
			"1",
			3,
			[][]int{{0, 1}, {1, 2}, {0, 2}},
			[]float64{0.5, 0.5, 0.2},
			0,
			2,
			0.25,
		},
		{
			"2",
			3,
			[][]int{{0, 1}, {1, 2}, {0, 2}},
			[]float64{0.5, 0.5, 0.3},
			0,
			2,
			0.3,
		},
		{
			"3",
			5,
			[][]int{{2, 3}, {1, 2}, {3, 4}, {1, 3}, {1, 4}, {0, 1}, {2, 4}, {0, 4}, {0, 2}},
			[]float64{0.06, 0.26, 0.49, 0.25, 0.2, 0.64, 0.23, 0.21, 0.77},
			0,
			3,
			0.16,
		},
	}

	for _, test := range testCases {
		t.Run(test.name, func(t *testing.T) {
			result := maxProbabilityHeap(test.n, test.edges, test.succProb, test.start, test.end)
			if result != test.expected {
				t.Fatalf("Test case #%s failed: expected: %f got: %f", test.name, test.expected, result)
			}
		})
	}
}

func maxProbabilityHeap(n int, edges [][]int, succProb []float64, start int, end int) float64 {
	// 1. build graph
	graph := make([][]node, n)
	for i := range edges {
		source := edges[i][0]
		dest := edges[i][1]
		graph[source] = append(graph[source], node{dest, succProb[i]})
		graph[dest] = append(graph[dest], node{source, succProb[i]})
	}

	visited := make([]bool, n)
	probs := make([]float64, n)
	probs[start] = 1

	gHeap := graphHeap{}
	gHeap = append(gHeap, node{start, 1})

	// 2. iterate graph
	for len(gHeap) > 0 {
		// 3. pop node with max probability
		source := heap.Pop(&gHeap).(node)
		if source.node == end {
			break
		}

		visited[source.node] = true

		// 4. calculate cost to each neighbor node
		for _, dest := range graph[source.node] {
			if visited[dest.node] {
				continue
			}

			prob := probs[source.node] * dest.prob
			if prob <= probs[dest.node] {
				continue
			}

			// 5. update dest node prob if it's bigger and push to heap
			dest.prob = prob
			probs[dest.node] = prob
			heap.Push(&gHeap, dest)
		}
	}

	return probs[end]
}

type graphHeap []node

func (h graphHeap) Len() int {
	return len(h)
}

// Less is inverted to creat MaxHeap instead of MinHeap
func (h graphHeap) Less(i, j int) bool {
	return h[i].prob > h[j].prob
}

func (h graphHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *graphHeap) Push(x any) {
	*h = append(*h, x.(node))
}

func (h *graphHeap) Pop() any {
	node := (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]
	return node
}

type node struct {
	node int
	prob float64
}
