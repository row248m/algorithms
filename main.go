package main

import (
	"fmt"
	"row248/algorithms/sort"
)

func main() {
	input := sort.GenerateInput()

	fmt.Println("input: ", input)
	fmt.Println("output: ", sort.MergeSort(input))
}
