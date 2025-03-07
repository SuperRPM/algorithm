package main

import (
	"fmt"
	"strconv"
)

func main() {
	inputEx := []string{
		"02984", "567",
	}

	for _, value := range inputEx {
		solution(value)
	}

}

func solution(numbers string) {
	result := 0
	for _, stringNumber := range numbers {
		number, _ := strconv.Atoi(string(stringNumber))
		if number == 0 || number == 1 || result == 0 {
			result += number
		} else {
			result *= number
		}
	}
	fmt.Println(result)
}
