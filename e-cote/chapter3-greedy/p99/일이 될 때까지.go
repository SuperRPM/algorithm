package main

import "fmt"

func main() {
	n := 25
	k := 5
	count := 0

	for {
		if n == 1 {
			fmt.Println(count)
			break
		}

		if n%k == 0 {
			n = n / k

		} else {
			n--
		}
		count++
	}
}
