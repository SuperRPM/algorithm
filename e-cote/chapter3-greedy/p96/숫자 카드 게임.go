package main

import (
	"fmt"
	"slices"
)

func main() {
	list1 := [][]int{
		{3, 1, 2},
		{4, 1, 4},
		{2, 2, 2},
	}

	list2 := [][]int{
		{7, 3, 1, 8},
		{3, 3, 3, 4},
	}

	solution(list1)
	solution(list2)
}

func solution(list [][]int) {
	var biggestInTheMinimis int
	for _, row := range list {
		if slices.Min(row) > biggestInTheMinimis {
			biggestInTheMinimis = slices.Min(row)
		}
	}
	fmt.Println(biggestInTheMinimis)
}
