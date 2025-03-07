package main

import (
	"fmt"
	"strconv"
)

func main() {
	ex := []string{
		"123402",
		"7755",
	}

	for _, n := range ex {
		solution(n)
	}
}

func solution(n string) {
	length := len(n) / 2

	left := n[:length]
	right := n[length:]

	leftSum := makeSum(left)
	rightSum := makeSum(right)
	if leftSum == rightSum {
		fmt.Println("LUCKY")
		return
	}
	fmt.Println("READY")

}

func makeSum(score string) int {
	result := 0
	for _, number := range score {
		toAdd, _ := strconv.Atoi(string(number))
		result += toAdd
	}
	return result
}
