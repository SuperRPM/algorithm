package main

import (
	"fmt"
	"sort"
)

func main() {
	input := [][]int{
		{5, 8, 3},
		{2, 4, 5, 4, 6},
	}

	// n := input[0][0]
	m := input[0][1]
	k := input[0][2]

	numberList := input[1]
	sort.Slice(numberList, func(i, j int) bool {
		return numberList[i] > numberList[j]
	})

	biggest := numberList[0]
	second := numberList[1]

	L := (m % k) * k

	result := biggest * L
	result += second * (m - L)

	fmt.Println(result)
}
