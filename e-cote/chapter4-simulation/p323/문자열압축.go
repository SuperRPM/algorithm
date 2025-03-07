package main

import (
	"fmt"
	"strconv"
)

func main() {
	ex := []string{
		// "aabbaccc",
		// "ababcdcdababcdcd",
		"abcabcdede",
		// "abcabcabcabcdededededede",
		// "xababcdcdababcdcd",
	}

	for _, s := range ex {
		solution(s)
	}
}

func solution(s string) {

	result := len(s)
	for gap := 1; gap < len(s)/2+1; gap++ {
		comp := ""
		prev := s[0:gap]
		count := 1
		for i := gap; i < len(s); i += gap {
			if i+gap > len(s) {
				strCount := strconv.Itoa(count)
				if count != 1 {
					comp += strCount
				}
				comp += prev
				comp += s[i:]
				prev = ""
				break
			}
			if prev == s[i:i+gap] {
				count++
			} else {
				strCount := strconv.Itoa(count)
				if count != 1 {
					comp += strCount
				}
				comp += prev
				prev = s[i : i+gap]
				count = 1
			}
		}
		strCount := strconv.Itoa(count)
		if count != 1 {
			comp += strCount
		}
		comp += prev
		if len(comp) < result {
			fmt.Println(comp)
			result = len(comp)
		}

	}
	fmt.Println(result)
}
