package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func main() {
	ex := []string{
		"K1KA5CB7",
		"AJKDLSI412K4JSJ9D",
	}

	for _, s := range ex {
		solution(s)
	}
}

func solution(input string) {
	charSlice := make([]string, 0)
	sum := 0
	for _, c := range input {
		intC, err := strconv.Atoi(string(c))
		if err != nil {
			charSlice = append(charSlice, string(c))
		} else {
			sum += intC
		}
	}

	sort.Slice(charSlice, func(i, j int) bool {
		return charSlice[i] < charSlice[j]
	})

	result := strings.Join(charSlice, "")

	strSum := strconv.Itoa(sum)
	result += strSum

	fmt.Println(result)
}
