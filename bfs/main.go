package main

import "fmt"

func main() {
	graph := [][]int{
		{1, 3},
		{2},
		{4, 7},
		{5, 6, 7},
		{6},
		{7},
	}

	fmt.Println(findShortestPath(graph, 0, 7, 8))
}

func findShortestPath(graph [][]int, src, dst, n int) []int {
	dist := make([]int, n)
	pred := make([]int, n)

	for i := 0; i < n; i++ {
		dist[i] = 1000000
		pred[i] = -1
	}

	if !bfs(graph, src, dst, n, pred, dist) {
		fmt.Println("not found path to dst")
		return nil
	}

	crawl := dst
	path := []int{crawl}

	for pred[crawl] != -1 {
		path = append(path, pred[crawl])
		crawl = pred[crawl]
	}

	return path
}

func bfs(graph [][]int, src, dst, n int, pred, dist []int) bool {
	queue := make([]int, 0, n)
	visited := make([]bool, n)

	visited[src] = true
	queue = append(queue, src)
	dist[src] = 0

	for i := 0; i < len(queue); i++ {
		for _, v := range graph[queue[i]] {
			if visited[v] {
				continue
			}

			visited[v] = true
			dist[v] = dist[queue[i]] + 1
			pred[v] = queue[i]
			queue = append(queue, v)

			if v == dst {
				return true
			}
		}
	}

	return false
}
