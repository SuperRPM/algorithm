package main

import (
	"fmt"
	"strconv"
)

func main() {
	where := "a1"

	column := convertColumn(string(where[0]))
	row, _ := strconv.Atoi(string(where[1]))
	row--

	solution(row, column)
}

func convertColumn(alphabet string) int {
	switch alphabet {
	case "a":
		return 0
	case "b":
		return 1
	case "c":

		return 2
	case "d":
		return 3
	case "e":
		return 4
	case "f":
		return 5
	case "g":
		return 6
	case "h":
		return 7
	}
	return 0
}

func solution(row, column int) {
	count := 0

	// up
	if row-2 >= 0 {
		if column+1 < 8 {
			count++
		}
		if column-1 >= 0 {
			count++
		}
	}

	// down
	if row+2 < 8 {
		if column+1 < 8 {
			count++
		}
		if column-1 >= 0 {
			count++
		}
	}

	// right
	if column+2 < 8 {
		if row+1 < 8 {
			count++
		}
		if row-1 >= 0 {
			count++
		}
	}

	// left
	if column-2 >= 0 {
		if row+1 < 8 {
			count++
		}
		if row-1 >= 0 {
			count++
		}
	}
	fmt.Println(count)
}
