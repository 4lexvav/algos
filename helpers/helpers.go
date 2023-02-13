package helpers

func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}

func MakeMatrix(m, n int) [][]int {
	matrix := make([][]int, m)
	for i := 0; i < m; i++ {
		matrix[i] = make([]int, n)
	}
	return matrix
}
