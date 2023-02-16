package main

import (
	"testing"
)

func TestWordSearch(t *testing.T) {
	tests := []struct {
		name     string
		board    [][]byte
		word     string
		expected bool
	}{
		{
			"1",
			[][]byte{{'C', 'A', 'A'}, {'A', 'A', 'A'}, {'B', 'C', 'D'}},
			"AAB",
			true,
		},
		{
			"2",
			[][]byte{{'A', 'B', 'C', 'E'}, {'S', 'F', 'E', 'S'}, {'A', 'D', 'E', 'E'}},
			"ABCESEEEFS",
			true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := exist(tt.board, tt.word)
			if result != tt.expected {
				t.Fatalf("Test case #%s failed, expected %v, got: %v", tt.name, tt.expected, result)
			}
		})
	}
}

func exist(board [][]byte, word string) bool {
	for i := range board {
		for j := range board[i] {
			if board[i][j] == word[0] {
				if dfs(board, word, i, j) {
					return true
				}
			}
		}
	}

	return false
}

func dfs(board [][]byte, word string, i, j int) bool {
	if len(word) == 0 {
		return true
	}

	if i >= len(board) || j >= len(board[0]) || i < 0 || j < 0 || board[i][j] != word[0] {
		return false
	}

	letter := board[i][j]
	board[i][j] = '#'

	if dfs(board, word[1:], i, j+1) {
		return true
	}

	if dfs(board, word[1:], i, j-1) {
		return true
	}

	if dfs(board, word[1:], i+1, j) {
		return true
	}

	if dfs(board, word[1:], i-1, j) {
		return true
	}

	// we set the letter back to the cell because we cannot move anywhere from it,
	// but we should be able to step into it again from another cell
	board[i][j] = letter

	return false
}

func dfsVisited(board [][]byte, word string, visited []bool, i, j int) bool {
	if len(word) == 0 {
		return true
	}

	rowsLen := len(board)
	colsLen := len(board[0])
	cellID := i*colsLen + j

	if i >= rowsLen || j >= colsLen || i < 0 || j < 0 || visited[cellID] || board[i][j] != word[0] {
		return false
	}

	visited[cellID] = true

	if dfsVisited(board, word[1:], visited, i, j+1) {
		return true
	}

	if dfsVisited(board, word[1:], visited, i, j-1) {
		return true
	}

	if dfsVisited(board, word[1:], visited, i+1, j) {
		return true
	}

	if dfsVisited(board, word[1:], visited, i-1, j) {
		return true
	}

	// we mark this cell as not visited because we cannot move anywhere from it,
	// but we should be able to step into it again from another cell
	visited[cellID] = false

	return false
}
