package main

import "fmt"

func main() {
	key := [][]int{
		{0, 0, 0},
		{1, 0, 0},
		{0, 1, 0},
	}
	lock := [][]int{
		{1, 1, 1},
		{1, 1, 0},
		{1, 0, 1},
	}

	answer := false
	answer = tryUnlock(key, lock)
	if answer {
		fmt.Println(answer)
		return
	}
	for i := 0; i < 3; i++ {
		key = rotate(key)
		answer = tryUnlock(key, lock)
		if answer {
			fmt.Println(answer)
			return
		}
	}
	fmt.Println(answer)

}

func rotate(key [][]int) [][]int {
	n := len(key)
	nkey := make([][]int, n)
	for i := 0; i < n; i++ {
		nkey[i] = make([]int, 0)
		for j := n - 1; j >= 0; j-- {
			nkey[i] = append(nkey[i], key[j][i])
		}
	}
	return nkey
}

func tryUnlock(key, lock [][]int) bool {
	n := len(lock)
	newLock := make([][]int, 3*n)
	for i := 0; i < 3*n; i++ {
		newRow := make([]int, 3*n)
		newLock[i] = append([]int(nil), newRow...)
	}

	// 가운데에 원래 lock 집어 넣기
	for i := n; i < 2*n; i++ {
		for j := n; j < 2*n; j++ {
			newLock[i][j] = lock[i-n][j-n]
		}
	}

	for i := 1; i < 2*n; i++ {
		for j := 1; j < 2*n; j++ {
			for x := 0; x < n; x++ {
				for y := 0; y < n; y++ {
					newLock[i+x][j+y] += key[x][y]
				}
			}
			result := true
			for i := n; i < 2*n; i++ {
				for j := n; j < 2*n; j++ {
					if newLock[i][j] != 1 {
						result = false
					}
				}
			}
			if !result {
				for x := 0; x < n; x++ {
					for y := 0; y < n; y++ {
						newLock[i+x][j+y] -= key[x][y]
					}
				}
			} else {
				return true
			}
		}
	}

	return false
}
