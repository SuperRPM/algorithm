package main

import "fmt"

func main() {
	graph := [][]int{
		{0, 0, 0, 0, 0, 1, 1, 1, 1, 0, 0, 0, 0, 0},
		{1, 1, 1, 1, 1, 1, 0, 1, 1, 1, 1, 1, 1, 0},
		{1, 1, 0, 1, 1, 1, 0, 1, 1, 0, 1, 1, 1, 0},
		{1, 1, 0, 1, 1, 1, 0, 1, 1, 0, 0, 0, 0, 0},
		{1, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1},
		{1, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 0, 0},
		{1, 1, 0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 1, 1},
		{0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 1, 1},
		{0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 0, 0, 0},
		{0, 0, 0, 1, 1, 1, 1, 1, 1, 1, 1, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 1, 0, 0, 0},
		{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 0, 0, 1, 1},
		{1, 1, 1, 0, 0, 0, 1, 1, 1, 1, 1, 1, 1, 1},
		{1, 1, 1, 0, 0, 0, 1, 1, 1, 1, 1, 1, 1, 1},
	}

	n := len(graph)
	m := len(graph[0])

	result := 0
	for i := range n {
		for j := range m {
			if dfs(i, j, n, m, graph) {
				result++
			}
		}
	}

	fmt.Println(result)
}

func dfs(x, y, n, m int, graph [][]int) bool {
	if x < 0 || x >= n || y < 0 || y >= m {
		return false
	}

	if graph[x][y] == 0 {
		graph[x][y] = 1
		dfs(x-1, y, n, m, graph)
		dfs(x, y-1, n, m, graph)
		dfs(x+1, y, n, m, graph)
		dfs(x, y+1, n, m, graph)
		return true
	} else {
		return false
	}

}
