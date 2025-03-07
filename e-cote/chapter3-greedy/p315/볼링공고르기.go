package main

import "fmt"

func main() {
	ex := [][]int{
		{1, 3, 2, 3, 2},
		{1, 5, 4, 3, 2, 4, 5, 2},
	}

	for _, arr := range ex {
		solution(arr)
	}
}

func solution(arr []int) {
	result := 0
	for i := 0; i < len(arr); i++ {
		for j := i + 1; j < len(arr); j++ {
			if arr[i] != arr[j] {
				result++
			}
		}
	}
	fmt.Println(result)
}
