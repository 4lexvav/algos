package helpers

func matrix(m, n int) [][]int {
	matrix := make([][]int, m)
	for i := 0; i < m; i++ {
		matrix[i] = make([]int, n)
	}
	return matrix
}

// Structures for graphs

// [][0] row idx
// [][1] col idx
// [][2] effort needed to move to this cell
type maxHeap [][]int

func (h maxHeap) Len() int {
	return len(h)
}

func (h maxHeap) Less(i, j int) bool {
	return h[i][2] < h[j][2]
}

func (h maxHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *maxHeap) Push(x interface{}) {
	*h = append(*h, x.([]int))
}

func (h *maxHeap) Pop() interface{} {
	col := (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]
	return col
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
