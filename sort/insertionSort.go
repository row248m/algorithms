package main

import (
	"fmt"
	"math/rand"
	"time"
)

func generateInput() []int {
	res := make([]int, 10)

	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 10; i++ {
		res[i] = rand.Intn(100)
	}

	return res
}

func sort(items []int) []int {
	for i:= 1; i < len(items); i++ {
		val := items[i]
		k := i - 1

		for k >= 0 && items[k] > val {
			items[k], items[k + 1] = items[k + 1], items[k]
			k--
		}
	}

	return items
}

func main() {
	input := generateInput()

	fmt.Println("input: ", input)
	fmt.Println("output: ", sort(input))
}