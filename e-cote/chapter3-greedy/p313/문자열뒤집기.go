package main

import "fmt"

func main() {
	example := "01010101010101010"

	countByChar := map[string]int{
		"0": 0,
		"1": 0,
	}

	prev := ""
	for i, number := range example {
		char := string(number)
		if prev == char {

		} else {
			prev = char
			countByChar[prev]++
		}

		if i == len(example)-1 {
			countByChar[char]++
		}
	}

	result := 1_000_001
	for _, v := range countByChar {
		if v < result {
			result = v
		}
	}
	fmt.Println(result)
}
