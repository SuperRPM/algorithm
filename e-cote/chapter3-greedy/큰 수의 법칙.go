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

	count := 0
	lastIndex := 0
	howManyUsed := 0
	result := 0

	for {
		if count == m {
			break
		}

		for i, v := range numberList {
			if lastIndex == i {
				if howManyUsed < k {
					result += v
					howManyUsed += 1
					count++
					break
				}
			} else {
				result += v
				lastIndex = i
				howManyUsed = 1
				count++
				break
			}
		}
	}
	fmt.Println(result)
}
