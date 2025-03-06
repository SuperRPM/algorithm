package main

import (
	"container/list"
	"fmt"
)

func main() {
	graph := [][]int{
		{1, 0, 1, 0, 1, 0},
		{1, 1, 1, 1, 1, 1},
		{0, 0, 0, 0, 0, 1},
		{1, 1, 1, 1, 1, 1},
		{1, 1, 1, 1, 1, 1},
	}

	x := 0
	y := 0
	bfs(x, y, graph)
}

func bfs(x, y int, graph [][]int) {
	q := list.New()
	q.PushBack([]int{x, y})

	// 방향 정의 (상, 하, 좌, 우)
	directions := []struct{ dx, dy int }{
		{-1, 0}, // 상
		{1, 0},  // 하
		{0, -1}, // 좌
		{0, 1},  // 우
	}
	result := 0
	for q.Len() > 0 {
		elem := q.Front()
		result++
		q.Remove(elem)
		coord := elem.Value.([]int)
		cx, cy := coord[0], coord[1]
		fmt.Println(cx, cy)
		if cx == len(graph)-1 && cy == len(graph[0])-1 {
			break
		}

		for _, direction := range directions {
			nx, ny := cx+direction.dx, cy+direction.dy
			if nx >= 0 && ny >= 0 && nx < len(graph) && ny < len(graph[0]) && graph[nx][ny] == 1 && graph[nx][ny] != 0 {
				graph[nx][ny] = graph[cx][cy] + 1
				q.PushBack([]int{nx, ny})
			}
		}
	}
	fmt.Println(graph[len(graph)-1][len(graph[0])-1])

}
