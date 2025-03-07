package main

import (
	"container/list"
	"fmt"
)

func main() {
	k := 5
	foodTimes := list.New()
	foodTimes.PushBack(3)
	foodTimes.PushBack(1)
	foodTimes.PushBack(2)
	lenfoodies := foodTimes.Len()
	result := 1
	for k > 0 {
		toEat := foodTimes.Front()
		foodTimes.Remove(toEat)

		toEatRemain, ok := toEat.Value.(int)
		if !ok {
			fmt.Println("not ok")
		}
		if toEatRemain == 0 {
			result++
			continue
		}
		toEatRemain--
		foodTimes.PushBack(toEatRemain)
		result++
		k -= 1
	}
	if foodTimes.Len() == 0 {
		fmt.Println(-1)
		return
	}

	fmt.Println(result % lenfoodies)

}
