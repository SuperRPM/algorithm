package main

import (
	"fmt"
	"slices"
)

// 못푼문제
func main() {

	example := []int{
		3, 2, 1, 1, 9,
	}

	slices.Sort(example)
	target := 1

	for _, number := range example {
		if target >= number {
			target += number
		} else {
			break
		}
	}

	fmt.Println(target)
}
