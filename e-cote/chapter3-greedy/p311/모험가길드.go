package main

import (
	"fmt"
	"slices"
)

func main() {
	// 순회하면서
	// len(arr)보다 크면 바로 continue
	// 1이면 그룹 수로 바로 카운트
	// 2이상이면 다음거를 포함해서
	// 요소의 max보다 len(arr)이 크면 즉시 카운트 + 1 및 arr 초기화

	arr := []int{2, 3, 1, 2, 2}
	count := 0
	group := make([]int, 0)
	for _, number := range arr {
		group = append(group, number)
		if slices.Max(group) <= len(group) {
			count++
			group = []int{}
		}
	}
	fmt.Println(count)
}
