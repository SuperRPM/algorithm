package main

import (
	"fmt"
	"slices"
)

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
